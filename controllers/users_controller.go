package controllers

import (
	"encoding/json"
	"go-microservices/services"
	"go-microservices/utils"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiError)
		w.WriteHeader(apiError.StatusCode)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonValue)
		return
	}

	user, apiError := services.GetUser(userId)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(apiError.StatusCode)
		w.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonValue)
}
