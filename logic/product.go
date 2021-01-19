package logic

import (
	"study/dao/mysql"
	"study/models"
)

func GetProductList(page,limit int64)(data []*models.Product,err error){
	data,err = mysql.GetProductList(page,limit)
	if err != nil {
		return
	}
	return
}
func GetProductByNameList(name string)(data []*models.Product,err error){
	productList,err := mysql.GetProductByNameList(name)
	if err != nil {
		return
	}
	data = make([]*models.Product,0,len(productList))
	data = productList
	return data,nil
}
