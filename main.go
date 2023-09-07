package main

import (
	"fmt"
	"log"
	"os"

	db "github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Server starting...")
	godotenv.Load()
	fmt.Println("Loaded env variables...")

	db.Connect()

	app := fiber.New()
	app.Use(cors.New())

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT is not set in .env file")
	}
	routes.Setup(app)
	app.Listen(":" + portString)
}
