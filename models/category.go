package models

type Category struct {
	Id int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	TopicImgId int `db:"topic_img_id" json:"topic_img_id"`
	*Image
}
type Image struct {
	Url string `db:"url" json:"url"`
}
