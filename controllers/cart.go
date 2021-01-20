package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"strconv"
	"study/logic"
	"study/models"
)

func AddCart(c *gin.Context){
	var cart *models.ParamCart
	if err:=c.ShouldBind(&cart);err != nil {
		zap.L().Error("Get Cart List Occur fail",zap.Error(err))
		errs,ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c,CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		return
	}
	user_id,err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}
	err = logic.AddCart(user_id,cart.ProductId,cart.Number)
	if err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}
	ResponseSuccess(c,CodeSuccess)

}
func DelCart(c *gin.Context){
	id := c.Query("id")
	ids, _ := strconv.Atoi(id)
	err := logic.DelCart(ids)
	if err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}
	ResponseSuccess(c,CodeSuccess)
}
func GetCart(c *gin.Context){
	user_id,err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c,CodeInvalidParams)
		return
	}
	cart,err := logic.GetCart(user_id)
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,err)
		return
	}
	ResponseSuccess(c,cart)
}
