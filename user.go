package trip

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	AccessToken     string `json:"access_token"`
	UserVkId int 	`json:"user_vk_id" db:"user_vk_id"`
	PhotoUrl string `json:"photo_url" db:"photo_url" binding:"required"`
}