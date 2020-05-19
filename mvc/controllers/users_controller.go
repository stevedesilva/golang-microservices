package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/stevedesilva/golang-microservices/mvc/services"
	"github.com/stevedesilva/golang-microservices/mvc/utils"
)

// GetUser func
func GetUser(resp http.ResponseWriter, req *http.Request) {
	userID, err := strconv.Atoi(req.URL.Query().Get("user_id"))
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	user, apiErr := services.GetUser(userID)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)

}
