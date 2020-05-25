package controllers

import (
	"net/http"
	"strconv"

	"github.com/stevedesilva/golang-microservices/mvc/services"

	"github.com/gin-gonic/gin"
	"github.com/stevedesilva/golang-microservices/mvc/utils"
)

// GetUser func
func GetUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	user, apiErr := services.UsersService.GetUser(userID)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	utils.Respond(c, http.StatusOK, user)

}
