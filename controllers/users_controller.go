package controllers

import (
	"WebServer/mvc/services"
	"WebServer/mvc/utils"
	"encoding/json"
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
		w.Write(jsonValue)
		return
	}

	user, apiError := services.GetUser(userId)
	if apiError != nil {
		jsonValue, _ := json.Marshal(apiError)
		w.WriteHeader(apiError.StatusCode)
		w.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	w.Write(jsonValue)
}
