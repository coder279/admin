package logic

import (
	"errors"
	"fmt"
	"study/dao/mysql"
	"study/models"
	"study/pkg/jwt"
	"study/pkg/snowflake"
)
var uuid uint64
func LogicSignUp(p *models.ParamSignup) (err error){
	is,err := mysql.CheckUserExist(p.Mobile)
	if err != nil {
		return err
	}
	if is {
		return errors.New("用户已经存在")
	}
	uuid,_ = snowflake.GetID()
	fmt.Println(uuid)
	user := &models.User{
		UUID: uuid,
		Nickname: p.Nickname,
		Password: p.Password,
		Mobile: p.Mobile,
		HeadImg: p.HeadImg,
	}
	err = mysql.InsertUser(user)
	if err != nil {
		return
	}
	return
}
func LogicSignIn(p *models.ParamLogin) (err error,token,refreshtoken string){
	is,err := mysql.CheckUserExist(p.Mobile)
	if err != nil {
		return err,"",""
	}
	if !is {
		return errors.New("用户不存在"),"",""
	}
	user := &models.User{
		Mobile:p.Mobile,
		Password:p.Password,
	}

	err,user_id := mysql.QueryUsersByMobile(user)
	if err != nil {
		return err,"",""
	}
	token,refreshtoken,err = jwt.GenToken(int64(user_id))
	if err != nil {
		return err,"",""
	}
	return nil,token,refreshtoken
}