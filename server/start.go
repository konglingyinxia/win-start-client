package server

import (
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/init/db"
	"github.com/konglingyinxia/win-start-client/server/init/log"
	"github.com/konglingyinxia/win-start-client/server/model"
	"os"
)

func Start() {
	log.Init()
	db.Init()
	initDb()
	initEnvKey()
}

func initDb() {
	err := global.DB.AutoMigrate(model.Env{}, model.App{}, model.GlobalSetting{})
	if err != nil {
		global.LOG.Error("sqlite-数据库表结构初始化失败...")
		return
	}
	global.LOG.Info("sqlite-数据库表结构初始化完成...")
}

// 初始化环境变量key
func initEnvKey() {
	//主应用目录
	appHomeValue := constant.GetApplicationHomePath()
	envHomePath := constant.GetEnvHomePath()
	err := os.Setenv(constant.EnvHomeKey, envHomePath)
	if err != nil {
		global.LOG.Error("初始化环境主目录变量失败...")
	}
	err = os.Setenv(constant.AppHomeKey, appHomeValue)
	if err != nil {
		global.LOG.Error("初始化应用主目录变量失败...")
	}
	global.LOG.Info("初始化[APP_HOME/ENV_HOME]环境变量完成...")
}
