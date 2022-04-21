package repository

import (
	"github.com/jmoiron/sqlx"
	trip "github.com/nvsces/flw-server-go"
)

type Authorization interface {
	CreateUser(user trip.User) (int, error)
	GetUser(user_vk_id int) (trip.User, error)
}


type TripItem interface {
	Create(item trip.TripItem) (int, error)
	GetAll() ([]ObjectOutputJson, error)
	GetById(itemId int) (trip.TripItem, error)
	Delete(userId,itemId int) error
}

type Profile interface{
	Info(userId int) (trip.User, error)
}

type Repository struct {
	Authorization
	TripItem
	Profile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TripItem: NewTripItemPostgres(db),
		Profile: NewProfilehPostgres(db),
	}
}