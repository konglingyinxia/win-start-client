package api

import (
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/dto/res"
	"github.com/konglingyinxia/win-start-client/server/utils/osutils"
)

// DashboardManager 概览
type DashboardManager struct {
}

func NewDashboardManager() *DashboardManager {
	return &DashboardManager{}
}

func (dm *DashboardManager) GetDashboard() (*res.OverviewData, error) {
	//获取应用状态
	apps, err := AppService.ListAll()
	if err != nil {
		return nil, err
	}
	status, err := AppService.RunningStatus(apps)
	if err != nil {
		return nil, err
	}
	var runningNum int
	for _, a := range status {
		if a.Status != constant.RuntimeStopped {
			runningNum++
		}
	}
	//运行环境数量
	envNum := EnvService.CountBy()
	//应用存储空间
	appPath := constant.GetApplicationHomePath()
	//环境存储空间
	envPath := constant.GetEnvHomePath()
	//日志存储空间
	logPath := constant.GetLogHomePath()
	return &res.OverviewData{
		RunningNum: runningNum,
		StoppedNum: len(status) - runningNum,
		EnvNum:     envNum,
		StorageApp: osutils.DirSize(appPath),
		StorageEnv: osutils.DirSize(envPath),
		StorageLog: osutils.DirSize(logPath),
	}, nil
}
