package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"study/logic"
	"study/models"
)
var productList []*models.Product

//商品列表
func ProductListHandler(c *gin.Context){
	var p models.ParamGetProductList
	if err := c.ShouldBind(&p);err != nil {
		zap.L().Error("Get Product List Occur fail",zap.Error(err))
		errs,ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c,CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		return
	}
	if p.Name == "" {
		fmt.Println(p.Page)
		productList,err := logic.GetProductList(int64(p.Page), int64(p.Limit))
		if err != nil {
			fmt.Println(err.Error())
			ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
			return
		}
		ResponseSuccess(c,productList)
	}else{
		productList,err := logic.GetProductByNameList(p.Name)
		if err != nil {
			fmt.Println(err.Error())
			ResponseErrorWithMsg(c,CodeServerBusy,err.Error())
			return
		}
		ResponseSuccess(c,productList)
	}
}
//商品详情
func ProductDetailHandler(c *gin.Context){

}
//商品分类
func ProductClassificationHandler(c *gin.Context){

}


