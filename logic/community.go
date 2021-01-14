package logic

import (
	"study/dao/mysql"
	"study/models"
)

func GetCommunityList()([]*models.Community,error){
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64)(*models.CommunityDetail,error){
	return mysql.GetCommunityDetailById(id)
}
