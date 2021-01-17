package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"study/models"
	"golang.org/x/crypto/bcrypt"
	"time"
)
//检测用户是否已经存在
func CheckUserExist(mobile string)(bool,error){
	query := `select count(id) from users where mobile = ?`
	var count int
	err := db.Get(&count,query,mobile)
	if err != nil {
		return false,err
	}
	return count>0,nil
}
//验证密码是否正确
func QueryUsersByMobile(user *models.User) (err error){
	password := user.Password
	query := `select id,mobile,password from users where mobile = ?`
	err = db.Get(user,query,user.Mobile)
	if err == sql.ErrNoRows{
		return errors.New("用户不能存在")
	}
	if err != nil {
		return err
	}
	res := checkpassword(password,user.Password)
	if !res {
		return errors.New("密码错误")
	}
	return nil

}
func checkpassword(password ,loginPwd string) (bool){
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginPwd)) //验证（对比
	if err != nil {
		return false
	}
	return true
}
//加密密码
func encrypassword(password string) (error,string){
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //加密处理
	if err != nil {
		fmt.Println(err)
	}
	encodePWD := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	return nil,encodePWD
}
//存储用户数据
func InsertUser(user *models.User) (err error){
	fmt.Println(user.Password)
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	err,password := encrypassword(user.Password)
	fmt.Println(password)
	query := `insert into users(uuid,nickname,head_img,mobile,password,login_time,created_at,updated_at) values(?,?,?,?,?,?,?,?)`
	db.Exec(query,user.UUID,user.Nickname,user.HeadImg,user.Mobile,password,currentTime,currentTime,currentTime)
	return
}
