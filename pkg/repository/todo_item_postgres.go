package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	trip "github.com/nvsces/flw-server-go"
)

type TripItemPostgres struct {
	db *sqlx.DB
}

func NewTripItemPostgres(db *sqlx.DB) *TripItemPostgres {
	return &TripItemPostgres{db: db}
}

func (r *TripItemPostgres) Create(item trip.TripItem) (int, error) {

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (author_id, date, route, count, type) values ($1, $2, $3, $4, $5) RETURNING id", tripItemsTable)

	row := r.db.QueryRow(createItemQuery, item.Author, item.Date, item.Route, item.Count, item.Type)
	if err :=row.Scan(&itemId); err!=nil {
		return 0, err
	}

	return itemId, nil
}

type ObjectOutputJson struct {
	User      	trip.User  	`json:"author" db:"author_id" binding:"required"`
	Date 		string `json:"date" db:"date"`
	Route 		string `json:"route" db:"route"`
	Count 		string `json:"count" db:"count"`
	Type 		string `json:"type" db:"type"`
	ItemId      int `json:"itemId"`
}

func (r *TripItemPostgres) GetAll() ([] ObjectOutputJson, error) {
	var items []trip.TripItemClient
	query := fmt.Sprintf(`select ti.id, ti.author_id, ti.date, ti.route, ti.count, ti.type, li.user_vk_id, li.name, li.photo_url from trip_items ti inner join users li on li.id = ti.author_id`)
	if err := r.db.Select(&items, query); err != nil {
		return nil, err
	}
	  v := []ObjectOutputJson{};
	for i:= 0; i < len(items); i++{
		var obj = trip.User{
			Id:items[i].AuthorId,
			Name: items[i].Name,
			PhotoUrl: items[i].PhotoUrl,
			UserVkId: items[i].UserVkId,
		};

		v = append(v,ObjectOutputJson{
			User: obj,
			ItemId: items[i].Id,
			Date: items[i].Date,
			Route: items[i].Route,
			Count: items[i].Count,
			Type: items[i].Type,
		} )
	}

	return v, nil
}
 					
func (r *TripItemPostgres) GetById(itemId int) (trip.TripItem, error) {
	var item trip.TripItem
	query := fmt.Sprintf(`SELECT * FROM trip_items WHERE trip_items.id = $1`)
	if err := r.db.Get(&item, query, itemId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TripItemPostgres) Delete(itemId int) error {
	query := fmt.Sprintf(`DELETE  FROM trip_items WHERE trip_items.id = $1`)
	_, err := r.db.Exec(query, itemId)
	return err
}
