package controllers

import (
	"github.com/gin-gonic/gin"
	"study/logic"
)

func GetCategory(c *gin.Context){
	categorys,err := logic.GetCategoryLogic()
	if err != nil {
		ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
		return
	}
	ResponseSuccess(c,categorys)
}