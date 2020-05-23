package services

import (
	"github.com/stevedesilva/golang-microservices/mvc/domain"
	"github.com/stevedesilva/golang-microservices/mvc/utils"
)

// UsersService var expose singleton
var UsersService usersService

type usersService struct{}

// GetUser func
func (u *usersService) GetUser(userID int) (*domain.User, *utils.ApplicationError) {

	user, err := domain.UserDao.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
