package services

import (
	"github.com/stevedesilva/golang-microservices/mvc/domain"
	"github.com/stevedesilva/golang-microservices/mvc/utils"
)

// GetUser func
func GetUser(userID int) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userID)
}
