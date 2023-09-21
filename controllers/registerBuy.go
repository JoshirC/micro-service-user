package controllers

import (
	db "github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/models"
	"github.com/gofiber/fiber/v2"
)

func CreateRegisterBuy(c *fiber.Ctx) error {
	var registerBuy models.RegisterBuy
	if err := c.BodyParser(&registerBuy); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Bad request",
			"error":   err,
		})
	}
	db.DB.Create(&registerBuy)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    registerBuy,
	})
}
func GetRegisterBuy(c *fiber.Ctx) error {
	var registerBuys []models.RegisterBuy

	db.DB.Select("name, rut, password, email, city, id_Buy, total_Buy, date_Buy").Find(&registerBuys)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    registerBuys,
	})
}
