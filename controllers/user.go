package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"study/dao/redis"
	"study/logic"
	"study/models"
	"study/pkg/sms"
	"study/settings"
)
var token = ""
var refreshToken = ""

func SignUpHandler(c *gin.Context){
	//1. 参数校验
	var p models.ParamSignup
	if err := c.ShouldBind(&p);err != nil {
		zap.L().Error("Signup with invalid param",zap.Error(err))
		errs,ok := err.(validator.ValidationErrors)
		if !ok{
			ResponseErrorWithMsg(c,CodeInvalidParams,errs)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		return
	}
	//2. 业务处理
	err := logic.LogicSignUp(&p)
	fmt.Println(err)
	if err != nil {
		ResponseErrorWithMsg(c,CodeUserExist,err.Error())
		return
	}
	//3. 返回响应
	ResponseSuccess(c,CodeSuccess)
}
func SignInHandler(c *gin.Context){
	var p models.ParamLogin
	if err := c.ShouldBind(&p);err != nil {
		zap.L().Error("Signin with invalid param",zap.Error(err))
		errs,ok := err.(validator.ValidationErrors)
		if !ok{
			ResponseErrorWithMsg(c,CodeInvalidParams,errs)
			return
		}
		ResponseErrorWithMsg(c,CodeInvalidParams,removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务处理
	if len(p.Code) == 0 {
		err,token,refreshToken := logic.LogicSignIn(&p)
		if err != nil {
			ResponseErrorWithMsg(c,CodeInvalidParams,err.Error())
			return
		}
		//3. 返回响应
		ResponseSuccess(c,gin.H{
			"token":token,
			"refreshToken":refreshToken,
		})
	}else{
		err,token,refreshToken := logic.LogicSignIn(&p)
		if err != nil {
			ResponseErrorWithMsg(c,CodeInvalidParams,err.Error())
			return
		}
		//3. 返回响应
		ResponseSuccess(c,gin.H{
			"token":token,
			"refreshToken":refreshToken,
		})
	}

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
	var m models.ParamMobile
	_= c.ShouldBind(&m)
	err := sms.Validation(m.Code,m.Mobile)
	if err != 1 {
		ResponseErrorWithMsg(c,CodeInvalidParams,"验证码错误或者过期")
		return
	}
	_ = redis.FogetCode(m.Mobile)

	ResponseSuccess(c,CodeSuccess)
}