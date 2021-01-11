package setting

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(configs ...string)(*Setting,error){
	vp := viper.New()
	vp.SetConfigName("config")
	for _,config := range configs {
		if config != ""{
			vp.AddConfigPath(config)
		}
	}
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil,err
	}
	s := &Setting{vp}
	s.WatchSettingConfig()
	return s,nil
}

func (s *Setting) WatchSettingConfig(){
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			s.ReadAllSection()
			fmt.Println("配置文件修改了")
		})
	}()
}