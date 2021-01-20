package models

type Cart struct {
	Id int `db:"id" json:"id"`
	Number int `db:"number" json:"number"`
	*Product
}
