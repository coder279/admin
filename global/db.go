package global

import (
	"lianquan/pkg/logger"
	"lianquan/pkg/setting"
)
var (
	ServerSetting *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
)
