package services_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevedesilva/golang-microservices/mvc/domain"
	"github.com/stevedesilva/golang-microservices/mvc/services"
	"github.com/stevedesilva/golang-microservices/mvc/utils"
)

var (
	userDaoMock usersDaoMock

	getUserFunction func(int) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userID int) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userID)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	getUserFunction = func(userId int) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{
			Message:    "user 0 was not found",
			StatusCode: http.StatusNotFound,
			Code:       "bad_request",
		}
	}

	user, err := services.UsersService.GetUser(0)
	assert.Nil(t, user)
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "user 0 was not found", err.Message)
}

func TestGetUserFound(t *testing.T) {
	getUserFunction = func(userId int) (*domain.User, *utils.ApplicationError) {
		return &domain.User{
			ID: 123,
		}, nil

	}

	user, err := services.UsersService.GetUser(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.ID)
}
