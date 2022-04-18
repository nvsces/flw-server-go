package repository

import (
	"fmt"

	trip "github.com/nvsces/flw-server-go"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user trip.User) (int, error) {
	
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, photo_url, user_vk_id) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.PhotoUrl, user.UserVkId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(user_vk_id int) (trip.User, error) {
	var user trip.User
	query := fmt.Sprintf("SELECT id, name, photo_url, user_vk_id  FROM %s WHERE user_vk_id=$1", usersTable)
	err := r.db.Get(&user, query, user_vk_id)

	return user, err
}