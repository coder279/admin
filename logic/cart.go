package logic

import (
	"study/dao/mysql"
	"study/models"
)

func AddCart(user_id int64,product_id,number int)(error){
	err := mysql.AddCart(user_id,product_id,number)
	return err
}
func GetCart(user_id int64)(cartList []*models.Cart,err error){
	return mysql.GetCart(user_id)
}
func DelCart(id int)(error){
	err := mysql.DelCart(id)
	return err
}
