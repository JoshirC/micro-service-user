package test

import (
	"testing"

	"github.com/JoshirC/micro-service-user.git/config"
	"github.com/JoshirC/micro-service-user.git/models"
)

func TestCreateUser(t *testing.T) {
	var reqCreateUser = models.Users{
		Name:     "juan",
		Rut:      "1234",
		Password: "12345",
		Email:    "juan@123",
		City:     "coquimbo",
	}
	result := config.DB.Create(&reqCreateUser)
	if result != nil {
		t.Error(result.Error)
	}
	t.Log("Insertado: ", result)
}
