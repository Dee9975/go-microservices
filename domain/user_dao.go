package domain

import (
	"WebServer/mvc/utils"
	"fmt"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 1, FirstName: "Fade", LastName: "Leon", Email: "fadedaf@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}
	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("User was not found %v", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
