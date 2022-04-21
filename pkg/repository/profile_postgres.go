package repository

import (
	"fmt"

	trip "github.com/nvsces/flw-server-go"

	"github.com/jmoiron/sqlx"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilehPostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}


func (r *ProfilePostgres) Info(userId int) (trip.User, error) {
	fmt.Println("postgrest")
	fmt.Println(userId);
	var user trip.User
	query := fmt.Sprintf("SELECT id, name, photo_url, user_vk_id  FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}