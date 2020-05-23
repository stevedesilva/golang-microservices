package domain_test

import (
	"net/http"
	"testing"

	"github.com/stevedesilva/golang-microservices/mvc/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetUserNotFound(t *testing.T) {
	user, err := domain.UserDao.GetUser(0)

	assert.Nil(t, user, "we are not expecting a user with id 0")
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := domain.UserDao.GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 1, user.ID)
	assert.EqualValues(t, "Max", user.FirstName)
	assert.EqualValues(t, "Payne", user.LastName)
	assert.EqualValues(t, "test@gmail.com", user.Email)

}
