package routes

import (
	"github.com/JoshirC/micro-service-user.git/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Wordl !")
	})

	//users routes
	app.Get("/users", controllers.GetUsers)
	app.Post("/users/new", controllers.CreateUser)
	//app.Delete("/users", controllers.DeleteUser)

	//registerBuy routes
	app.Get("/r", controllers.GetRegisterBuy)
	app.Post("/r", controllers.CreateRegisterBuy)

	//Login routes
	app.Post("/login", controllers.Login)
	app.Post("/singup", controllers.SingUp)

}
