package internal

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/JoshirC/micro-service-user.git/controllers"
	"github.com/JoshirC/micro-service-user.git/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Handler(d amqp.Delivery, ch *amqp.Channel) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var response models.Response
	var tokenString string

	var message struct {
		Pattern string          `json:"pattern"`
		Data    json.RawMessage `json:"data"`
		ID      string          `json:"id"`
	}

	err := json.Unmarshal(d.Body, &message)
	if err != nil {
		log.Printf("Error al desglosar el cuerpo del mensaje: %s", err)
		return
	}

	actionType := message.Pattern
	log.Printf(" [x] Received %s", actionType)
	switch actionType {
	case "LOGIN_USER":
		log.Println(" [.] Login user")

		var err error

		tokenString, err = controllers.Login(message.Data)

		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Login failed",
				Data:    []byte(err.Error()),
			}
		} else {
			user, err := controllers.GetUserByEmail(message.Data)
			JSONUser, err := json.Marshal(user)
			failOnError(err, "Failed to marshal user")
			response = models.Response{
				Success: "success",
				Message: "User logged",
				Data:    JSONUser,
			}
		}

	case "SIGNUP_USER":
		log.Println(" [.] Creating User")

		err := controllers.SingUp(message.Data)
		failOnError(err, "Failed to create User")

		response = models.Response{
			Success: "success",
			Message: "User Created",
			Data:    nil,
		}
	case "GET_USER":
		var data struct {
			ID uint `json:"id"`
		}

		log.Println(" [.] Getting user")
		err := json.Unmarshal(message.Data, &data)
		failOnError(err, "Failed to Unmarshal user")

		user, err := controllers.GetUser(data.ID)
		failOnError(err, "Failed to get user")
		userJSON, err := json.Marshal(user)
		failOnError(err, "Failed to marshal user")

		response = models.Response{
			Success: "success",
			Message: "User retrieved",
			Data:    userJSON,
		}

	case "CREATE_USER":
		var newUser models.Users
		err := json.Unmarshal(message.Data, &newUser)
		failOnError(err, "Failed to Unmarshal user")

		err = controllers.CreateUser(newUser)
		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Failed to create User",
				Data:    []byte(err.Error()),
			}
		} else {
			response = models.Response{
				Success: "success",
				Message: "User Created",
				Data:    nil,
			}
		}

	case "UPDATE_USER":
		log.Println(" [.] Updating User")

		var updatedUser models.Users
		err := json.Unmarshal(message.Data, &updatedUser)
		failOnError(err, "Failed to Unmarshal user")

		err = controllers.UpdateUser(updatedUser.ID, updatedUser)
		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Failed to update User",
				Data:    []byte(err.Error()),
			}
		} else {
			response = models.Response{
				Success: "success",
				Message: "User Updated",
				Data:    nil,
			}
		}

	case "DELETE_USER":
		log.Println(" [.] Deleting User")

		var userID struct {
			ID uint `json:"id"`
		}
		err := json.Unmarshal(message.Data, &userID)
		failOnError(err, "Failed to Unmarshal user")

		err = controllers.DeleteUser(userID.ID)
		if err != nil {
			response = models.Response{
				Success: "error",
				Message: "Failed to delete User",
				Data:    []byte(err.Error()),
			}
		} else {
			response = models.Response{
				Success: "success",
				Message: "User Deleted",
				Data:    nil,
			}
		}
	}

	responseJSON, err := json.Marshal(response)
	failOnError(err, "Failed to marshal response")

	err = ch.PublishWithContext(ctx,
		"",        // exchange
		d.ReplyTo, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: d.CorrelationId,
			Body:          responseJSON,
			Headers:       amqp.Table{"Authorization": tokenString},
		})
	failOnError(err, "Failed to publish a message")

	d.Ack(false)
}
