package services

import (
	"go-microservices/domain"
	"go-microservices/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
