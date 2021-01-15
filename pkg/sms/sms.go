package sms

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

func SendMsg(tel string, code string) string {
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "<accesskeyId>", "<accessSecret>")
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = tel //手机号变量值
	request.SignName = "益计划" //签名
	request.TemplateCode = "SMS_19586XXXX" //模板编码
	request.TemplateParam = "{\"code\":\"" + code + "\"}"
	response, err := client.SendSms(request)
	fmt.Println(response.Code)
	if response.Code == "isv.BUSINESS_LIMIT_CONTROL" {
		return "frequency_limit"
	}
	if err != nil {
		fmt.Print(err.Error())
		return "failed"
	}
	return "success"
}