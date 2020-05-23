package domain

import (
	"fmt"
	"net/http"

	"github.com/stevedesilva/golang-microservices/mvc/utils"
)

var (
	users = map[int]*User{
		123: {ID: 1, FirstName: "Max", LastName: "Payne", Email: "test@gmail.com"},
	}
	// UserDao used as singleton. Interface implements same as test
	// Public var 
	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

// private interface but public Method
type userDaoInterface interface {
	GetUser(int) (*User, *utils.ApplicationError)
}

type userDao struct{}

// GetUser func satisfies userDaoInterface
func (u *userDao) GetUser(userID int) (*User, *utils.ApplicationError) {
	if user := users[userID]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userID),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}

}
