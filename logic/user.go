package logic

import (
	"errors"
	"fmt"
	"study/dao/mysql"
	"study/models"
)
var uuid int64
func LogicSignUp(p *models.ParamSignup) (err error){
	is,err := mysql.CheckUserExist(p.Mobile)
	if err != nil {
		return err
	}
	if is {
		return errors.New("用户已经存在")
	}
	uuid = 1
	fmt.Printf("生成:%v",uuid)
	fmt.Println("进入这里1")
	user := &models.User{
		UUID: uuid,
		Nickname: p.Nickname,
		Password: p.Password,
		Mobile: p.Mobile,
		HeadImg: p.HeadImg,
	}
	fmt.Println("进入这里2")
	err = mysql.InsertUser(user)
	if err != nil {
		return
	}
	return
}