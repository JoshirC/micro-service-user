package controllers

import (
	db "github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/models"
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	var newUser models.Users

	if err := c.BodyParser(&newUser); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err,
		})
	}

	db.DB.Create(&newUser)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"data":    newUser,
	})
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.Users
	db.DB.Select("id, name, rut, password, email, city").Find(&users)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"data":    users,
	})
}

/* func DeleteUser(c *fiber.Ctx) error {
	userRut := c.Params("rut")
	var user models.Users
	db.DB.Select("name, rut, password, email, city").Where("rut = ?", userRut).First(&user)

	if user.Rut == "null" {
		return c.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	db.DB.Delete(&user)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
	})
} */
