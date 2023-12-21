package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/internal"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func getChannel() *amqp.Channel {
	ch := config.GetChannel()
	if ch == nil {
		log.Panic("Failed to get channel")
	}
	return ch
}

func declareQueue(ch *amqp.Channel) amqp.Queue {
	q, err := ch.QueueDeclare(
		"user_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare a queue")
	return q
}

func setQoS(ch *amqp.Channel) {
	err := ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")
}

func registerConsumer(ch *amqp.Channel, q amqp.Queue) <-chan amqp.Delivery {
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")
	return msgs
}

func init() {
	command := "docker-compose"
	args := []string{"up", "-d"}
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error al ejecutar el comando:", err)
		os.Exit(1)
	}

}

func main() {
	fmt.Println("Starting...")

	fmt.Println("Users MS starting...")

	godotenv.Load()
	fmt.Println("Loaded env variables...")

	config.SetupDatabase()
	fmt.Println("Database connection configured...")

	config.SetupRabbitMQ()
	fmt.Println("RabbitMQ Connection configured...")

	ch := getChannel()
	q := declareQueue(ch)
	setQoS(ch)
	msgs := registerConsumer(ch, q)

	var forever chan struct{}
	go func() {
		for d := range msgs {
			internal.Handler(d, ch)
		}
	}()
	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}
