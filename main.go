package main

import (
	"fmt"
	"go.uber.org/zap"
	"study/controllers"
	"study/dao/mysql"
	"study/dao/redis"
	"study/logger"
	"study/settings"
)
func init(){
	err := LoadingSetting()
	if err != nil {
		zap.L().Fatal("LoadingSetting err: ",zap.Error(err))
	}
}

func main() {


}

func LoadingSetting() error {
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed,%#v\n", err)
		return err
	}
	return nil
}
func LoadingLogger() error {
	if err := logger.Init(); err != nil {
		fmt.Printf("init settings failed,%#v\n", err)
		return err
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success")
	return nil
}
func LoadingMysql() error{
	if err := mysql.Init(settings.Conf.MySQLConfig); err != nil {
		fmt.Printf("mysql init settings failed,%#v\n", err)
		return err
	}
	defer mysql.Close()
	return nil
}
func LoadingRedis() error {
	if err := redis.Init(settings.Conf.RedisConfig); err != nil {
		fmt.Printf("redis init settings failed,%#v\n", err)
		return err
	}
	defer redis.Close()
	return nil
}
func LoadingValidator() error {
	if err := controllers.InitTrans("zh"); err != nil {
		fmt.Printf("init init trans failed,err:%v\n", err)
		return err
	}
	return nil
}
