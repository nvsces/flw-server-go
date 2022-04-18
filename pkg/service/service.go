package service

import (
	trip "github.com/nvsces/flw-server-go"
	"github.com/nvsces/flw-server-go/pkg/repository"
)

type Authorization interface {
	CreateUser(user trip.User) (int, error)
	GenerateToken(user_vk_id int) (string, error)
	ParseToken(token string) (int, error)
	GetUser(user_vk_id int) (trip.User, error)
}


type TripItem interface {
	Create(item trip.TripItem) (int, error)
	GetAll() ([]repository.ObjectOutputJson, error)
	GetById(itemId int) (trip.TripItem, error)
	Delete(userId,itemId int) error
}

type Service struct {
	Authorization
	TripItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TripItem: NewTripItemService(repos.TripItem),
	}
}