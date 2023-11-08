package controllers

import (
	"encoding/json"

	db "github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/models"
)

func GetUsers() ([]models.Users, error) {
	var users []models.Users
	err := db.DB.Select("id, first_name, last_name, email").Find(&users).Error
	return users, err
}

func GetUser(userkID uint) (models.Users, error) {
	var user models.Users
	err := db.DB.Select("id, title, description, done").Where("id = ?", userkID).First(&user).Error

	return user, err
}

func GetUserByEmail(body []byte) (models.Users, error) {
	var loginData models.LoginData

	err := json.Unmarshal(body, &loginData)
	if err != nil {
		return models.Users{}, err
	}

	var user models.Users
	err = db.DB.Select("id, first_name, last_name, email").Where("email = ?", loginData.Email).First(&user).Error
	return user, err
}
