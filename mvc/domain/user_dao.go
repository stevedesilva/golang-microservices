package domain

import (
	"fmt"
	"net/http"

	"github.com/stevedesilva/golang-microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {ID: 1, FirstName: "Steve", LastName: "deSilva", Email: "test@gmail.com"},
	}
)

// GetUser func
func GetUser(userID int64) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}

}
