package global

import (
	"github.com/julianlee107/blogWithGin/pkg/setting"
	"github.com/julianlee107/blogWithGin/pkg/logger"
)

var (
	ServerSetting   *setting.ServerSetting
	AppSetting      *setting.AppSetting
	DatabaseSetting *setting.DatabaseSetting
	Logger          *logger.Logger
)
