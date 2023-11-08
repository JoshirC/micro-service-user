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

func CreateUser(user models.Users) error {
	err := db.DB.Create(&user).Error
	return err
}

func DeleteUser(userID uint) error {
	err := db.DB.Where("id = ?", userID).Delete(&models.Users{}).Error
	return err
}

func UpdateUser(userID uint, updatedUser models.Users) error {
	err := db.DB.Model(&models.Users{}).Where("id = ?", userID).Updates(updatedUser).Error
	return err
}
