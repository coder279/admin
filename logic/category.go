package logic

import (
	"study/dao/mysql"
	"study/models"
)

func GetCategoryLogic()([]*models.Category,error){
	return mysql.GetCategory()
}
