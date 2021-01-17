package models

type User struct {
	UUID int64 `db:"uuid"`
	Nickname string `db:"nickname"`
	HeadImg string `db:"head_img"`
	Mobile string `db:"mobile"`
	Password string `db:"password"`
	LoginTime string `db:"login_time"`
	Token string
	*Model
}