package services

import (
	"WebServer/mvc/domain"
	"WebServer/mvc/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
