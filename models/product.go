package models

type Product struct {
	Id int `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Price string `json:"price "db:"price"`
	Stock string `json:"stock "db:"stock"`
	CategoryId int `json:"category_id "db:"category_id"`
	MainImgUrl string `json:"main_img_url "db:"main_img_url"`
	*ProductImg
}
type ProductImg struct {
	Url string `json:"url "db:"url"`
}

