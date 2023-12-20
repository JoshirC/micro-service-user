package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	db "github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(body []byte) (string, error) {
	var loginData models.LoginData

	err := json.Unmarshal(body, &loginData)

	if err != nil {
		return "", err
	}

	var user models.Users
	log.Printf("Comienzo busqueda: %s", loginData.Email)
	db.DB.Where("email = ?", loginData.Email).First(&user)
	log.Printf("Comienzo comparaciones: %v", user.ID)
	if user.ID == 0 {
		log.Printf("a: %s", err)
		return "", errors.New("User not found")
	}
	log.Printf("Comparacion password: %v", user.Password)
	log.Printf("Password entrada: %v", loginData.Password)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		log.Printf("aa: %s", err)
		return "", errors.New("Incorrect password")

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    user.ID,
		"ExpiresAt": time.Now().Add(time.Hour * 24 * 30).Unix(), //30 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		log.Printf("aaa: %s", err)
		return "", errors.New("Token Expired or invalid")
	}

	log.Printf("aaaaa: %s", tokenString)
	log.Printf("aaaaa: %s", err)
	return tokenString, err

}

func SingUp(body []byte) error {

	var singUpData models.SingUpData

	err := json.Unmarshal(body, &singUpData)
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(singUpData.Password), bcrypt.DefaultCost)

	if err != nil {
		return errors.New("Failed to hash password")
	}

	user := models.Users{
		Name:     singUpData.Name,
		Rut:      singUpData.Rut,
		Password: string(hash),
		Email:    singUpData.Email,
		City:     singUpData.City,
	}
	result := db.DB.Create(&user)

	if result.Error != nil {
		return errors.New("Failed to create user")
	}

	fmt.Println("Se agrego al usuario")
	return nil

}
