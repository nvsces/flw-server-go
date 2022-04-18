package trip

type TripItem struct {
	Id          int    `json:"id" db:"id"`
	Author      int   `json:"author" db:"author_id" binding:"required"`
	Date 		string `json:"date" db:"date"`
	Route 		string `json:"route" db:"route"`
	Count 		string `json:"count" db:"count"`
	Type 		string `json:"type" db:"type"`
}

type TripItemClient struct{
	Id          int    `json:"id" db:"id"`
	AuthorId      	int  	`json:"author_id" db:"author_id" binding:"required"`
	Date 		string `json:"date" db:"date"`
	Route 		string `json:"route" db:"route"`
	Count 		string `json:"count" db:"count"`
	Type 		string `json:"type" db:"type"`
	Name     	string `json:"name" binding:"required"`
	UserVkId int 	`json:"user_vk_id" db:"user_vk_id"`
	PhotoUrl string `json:"photo_url" db:"photo_url" binding:"required"`
}

