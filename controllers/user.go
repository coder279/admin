package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"study/dao/redis"
	"study/pkg/sms"
	"study/settings"
)

func SignUpHandler(c *gin.Context){
	//1. 参数校验
	var p ParamSignup
	fmt.Println(p)
	if err := c.ShouldBind(&p);err != nil {
		zap.L().Error("Signup with invalid param",zap.Error(err))
		errs,ok := err.(validator.ValidationErrors)
		if !ok{
			ResponseError(c,CodeInvalidParams)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		return
	}
	//2. 业务处理

	//3. 返回响应
	ResponseSuccess(c,CodeSuccess)
}

func SendSms(c *gin.Context)  {
	tel := c.Query("mobile")
	err := sms.SendMsg(settings.Conf.SmsConfig,tel,sms.Code())
	if err != nil {
		ResponseErrorWithMsg(c,CodeInvalidParams,err)
		return
	}
	ResponseSuccess(c,CodeSuccess)
}

func CheckSmsCodeValid(c *gin.Context){
	var m ParamMobile
	_= c.ShouldBind(&m)
	err := sms.Validation(m.Code,m.Mobile)
	if err != 1 {
		ResponseErrorWithMsg(c,CodeInvalidParams,"验证码错误或者过期")
		return
	}
	_ = redis.FogetCode(m.Mobile)

	ResponseSuccess(c,CodeSuccess)
}