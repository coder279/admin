package sms

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"study/dao/redis"
	"study/settings"
	"time"
)

func SendMsg(cfg *settings.SmsConfig,tel string, code string) error {
	client, err := dysmsapi.NewClientWithAccessKey(cfg.Region, cfg.AliossAccessId, cfg.AliossAcessKey)
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = tel //手机号变量值
	request.SignName = cfg.SignName //签名
	request.TemplateCode = cfg.TemplateCode //模板编码
	request.TemplateParam = "{\"code\":\"" + code + "\"}"
	response, err := client.SendSms(request)
	if response.Code == "isv.BUSINESS_LIMIT_CONTROL" {
		return err
	}
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	err = redis.SaveCode(tel,code)
	if err != nil {
		fmt.Print(err.Error())
		return err
	}
	return nil
}
// 随机验证码
func Code() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(899999) + 100000
	res := strconv.Itoa(code) //转字符串返回
	return res
}
// 在注册时检查验证码
func Validation(validation string, mobile string) int {
	fmt.Printf("手机号:%v",mobile)
	var flag int
	err, code := redis.GetCode(mobile)

	if err != nil {
		zap.L().Error("发送验证码出现错误",zap.Error(err))
	}
	if validation == code {
		flag = 1
	} else {
		flag = 0
	}
	return flag
}


