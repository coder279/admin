package mysql

import (
	"fmt"
	"study/models"
)

func GetProductList(page,limit int64)(product []*models.Product,err error){
	if(page == 1){
		page = 0
	}
	query := `select product.id as id,name,stock,category_id,main_img_url,price,url from product left join image on product.img_id = image.id Limit ?,?`
	err = db.Select(&product,query,page,limit)
	return
}
func GetProductByNameList(name string)(product []*models.Product,err error){
	query := `select product.id,name,stock,category_id,main_img_url,price,url from product left join image on product.img_id = image.id where name like ? `
	product = make([]*models.Product,0,2)
	err = db.Select(&product,query,name+"%")
	fmt.Println(err)
	return product,err
}

func GetProductDetail(id int)(product *models.Product,err error){
	query := `select product.id,name,stock,category_id,main_img_url,price,url from product left join image on product.img_id = image.id where id = ? limit 1`
	product = new(models.Product)
	err = db.Get(product,query,id)
	return product,err
}
func GetProductCategory(page,limit,category_id int)(product []*models.Product,err error){
	query := `select product.id,name,stock,category_id,main_img_url,price,url from product left join image on product.img_id = image.id where category_id = ? limit ?,?`
	product = make([]*models.Product,0,2)
	err = db.Select(product,query,category_id,page,limit)
	return product,err
}
func GetCategory()(category []*models.Category,err error){
	query := `select category.id,name,topic_img_id from category left join image on category.topic_img_id = image.id`
	category = make([]*models.Category,0,2)
	err = db.Select(category,query)
	return category,err
}