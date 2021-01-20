package redis

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func GetCode(key string) (error,string) {
	code,err := rdb.Get(key).Result()
	fmt.Println(code)
	if err != nil {
		return err,""
	}
	return nil,code
}
func SaveCode(key string,code string) error {
	err := rdb.Set(key,code,300 * time.Second)
	if err != nil {
		return err.Err()
	}
	return nil
}
func FogetCode(key string) error {
	if err := rdb.Del(key); err != nil {
		zap.L().Error("FogetCode出现错误:%v",zap.Error(err.Err()))
	}
	return nil
}
