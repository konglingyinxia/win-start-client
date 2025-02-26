package api

import (
	"github.com/konglingyinxia/win-start-client/server/dto/res"
)

type LogManager struct {
}

func NewLogManager() *LogManager {
	return &LogManager{}
}

// SystemLog Log 系统日志
func (a *LogManager) SystemLog() res.BaseRes {
	LogService.SystemLog()
	return res.Ok()
}

// CloseSystemLog 关闭系统日志
func (a *LogManager) CloseSystemLog() res.BaseRes {
	LogService.CloseSystemLog()
	return res.Ok()
}

// AppLog 应用日志
func (a *LogManager) AppLog(id string) res.BaseRes {
	app, err := AppService.GetById(id)
	if err != nil {
		return res.Err("获取应用失败:" + err.Error())
	}
	LogService.AppLog(app)
	return res.Ok()
}

// CloseAppLog 关闭APP日志
func (a *LogManager) CloseAppLog(id string) res.BaseRes {
	LogService.CloseAppLog(id)
	return res.Ok()
}
