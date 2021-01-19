package models

type Category struct {
	Id int `db:"id"`
	Name string `db:"name"`
	TopicImgId int `db:"topic_img_id"`
	Url string `db:"url"`
	*Image
}
type Image struct {
	Url string `db:"url"`
}
