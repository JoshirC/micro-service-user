package controllers_test

import (
	"encoding/json"
	"testing"

	"github.com/JoshirC/micro-service-user.git/controllers"
	"github.com/JoshirC/micro-service-user.git/models"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	usersToTest := []models.Users{
		{Name: "User1", Rut: "123", Email: "user1@example.com", City: "City1"},
		{Name: "User2", Rut: "456", Email: "user2@example.com", City: "City2"},
	}

	for _, user := range usersToTest {
		err := controllers.CreateUser(user)
		assert.NoError(t, err)
	}

	users, err := controllers.GetUsers()
	assert.NoError(t, err)

	assert.Len(t, users, len(usersToTest))

	for i, user := range users {
		assert.Equal(t, usersToTest[i].ID, user.ID)
		assert.Equal(t, usersToTest[i].Name, user.Name)
	}
}

func TestGetUser(t *testing.T) {
	userToTest := models.Users{Name: "TestUser", Rut: "789", Email: "test@example.com", City: "TestCity"}
	err := controllers.CreateUser(userToTest)
	assert.NoError(t, err)
	user, err := controllers.GetUser(userToTest.ID)
	assert.NoError(t, err)
	assert.Equal(t, userToTest.ID, user.ID)
	assert.Equal(t, userToTest.Name, user.Name)
}

func TestGetUserByEmail(t *testing.T) {
	userToTest := models.Users{Name: "TestUser", Rut: "789", Email: "test@example.com", City: "TestCity"}
	err := controllers.CreateUser(userToTest)
	assert.NoError(t, err)
	loginData := models.LoginData{Email: userToTest.Email}
	loginDataBytes, err := json.Marshal(loginData)
	assert.NoError(t, err)

	user, err := controllers.GetUserByEmail(loginDataBytes)
	assert.NoError(t, err)

	assert.Equal(t, userToTest.ID, user.ID)
	assert.Equal(t, userToTest.Name, user.Name)

}

func TestCreateUser(t *testing.T) {
	newUser := models.Users{Name: "John Doe", Rut: "123", Email: "john@example.com", City: "City1"}
	err := controllers.CreateUser(newUser)
	assert.NoError(t, err)

	createdUser, err := controllers.GetUserByEmail([]byte(`{"email": "john@example.com"}`))
	assert.NoError(t, err)
	assert.Equal(t, newUser.Name, createdUser.Name)
}

func TestDeleteUser(t *testing.T) {
	userToTest := models.Users{Name: "TestUser", Rut: "789", Email: "test@example.com", City: "TestCity"}
	err := controllers.CreateUser(userToTest)
	assert.NoError(t, err)
	err = controllers.DeleteUser(userToTest.ID)
	assert.NoError(t, err)
	_, err = controllers.GetUser(userToTest.ID)
	assert.Error(t, err)
}

func TestUpdateUser(t *testing.T) {
	userToTest := models.Users{Name: "TestUser", Rut: "789", Email: "test@example.com", City: "TestCity"}
	err := controllers.CreateUser(userToTest)
	assert.NoError(t, err)
	updatedUser := models.Users{Name: "UpdatedUser", Rut: "789", Email: "updated@example.com", City: "UpdatedCity"}
	err = controllers.UpdateUser(userToTest.ID, updatedUser)
	assert.NoError(t, err)
	user, err := controllers.GetUser(userToTest.ID)
	assert.NoError(t, err)
	assert.Equal(t, updatedUser.Name, user.Name)
	assert.Equal(t, updatedUser.Email, user.Email)
}
