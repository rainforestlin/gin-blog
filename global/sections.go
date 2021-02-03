package global

import (
	"github.com/julianlee107/blogWithGin/pkg/logger"
	"github.com/julianlee107/blogWithGin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger          *logger.Logger
	JWTSetting      *setting.JWTSetting
)
