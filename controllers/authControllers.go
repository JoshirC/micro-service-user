package controllers

import (
	"encoding/json"
	"errors"
	"os"
	"time"

	db "github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(body []byte) (string, error) {
	//Obtener el correo y contraseña desde el body del request
	var loginData models.LoginData

	err := json.Unmarshal(body, &loginData)
	if err != nil {
		return "", err
	}
	//Buscar al usuario en la DB
	var user models.Users
	db.DB.Where("email = ?", loginData.Email).First(&user)

	if user.ID == 0 {
		return "", errors.New("User not found")
	}

	//Verificar la contraseña

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password)); err != nil {
		return "", errors.New("Incorrect password")

	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    user.ID,
		"ExpiresAt": time.Now().Add(time.Hour * 24 * 30).Unix(), //30 days
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", errors.New("Token Expired or invalid")
	}

	return tokenString, nil
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

	//Respuesta exitosa
	return nil
}
