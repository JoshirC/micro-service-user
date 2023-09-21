package controllers

import (
	"net/http"

	db "github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.BodyParser(&body) != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}
	var user models.Users
	db.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"message": "Logged in",
		"data":    user,
	})
}
func SingUp(c *fiber.Ctx) error {
	var body struct {
		Name     string `json:"name"`
		Rut      string `json:"rut"`
		Password string `json:"password"`
		Email    string `json:"email"`
		City     string `json:"city"`
	}
	if c.BodyParser(&body) != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal serveral error",
		})
	}

	user := models.Users{
		Name:     body.Name,
		Rut:      body.Rut,
		Password: string(hash),
		Email:    body.Email,
		City:     body.City,
	}
	result := db.DB.Create(&user)

	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal error",
		})
	}

	return c.Status(http.StatusCreated).JSON(user)
}
