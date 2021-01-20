package mysql

import (
	"study/models"
	"time"
)

func AddCart(uuid int64,product_id,number int)(error){
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	query := `insert into cart(user_id,product_id,number,create_time,update_time) values(?,?,?,?,?)`
	_,err := db.Exec(query,int(uuid),product_id,number,currentTime,currentTime)
	return err
}
func GetCart(uuid int64)(cartList []*models.Cart,err error){
	cartList = make([]*models.Cart,0,2)
	query := `select cart.id,name,main_img_url,number from cart left join product on cart.product_id = product.id  where user_id = ? and cart.delete_time is NULL `
	err = db.Select(&cartList,query,int(uuid))
	return
}
func DelCart(id int)(error){
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	query := `update cart set delete_time = ? where id = ?`
	_, err:= db.Exec(query,currentTime,id)
	return err
}
