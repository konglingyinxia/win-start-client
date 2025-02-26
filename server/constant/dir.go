package constant

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/konglingyinxia/win-start-client/server/global"
	"os"
	"path/filepath"
)

const (
	ApplicationPath = "opt" + string(filepath.Separator) + "app"           // ApplicationPath 应用放置目录
	EnvPath         = "opt" + string(filepath.Separator) + "env"           // EnvPath 环境目录
	LogFilePath     = "logs" + string(filepath.Separator) + "catalina.log" // LogFilePath 日志文件
	DbFile          = "runtime.dt"                                         // DbFile 数据库文件
	ProcessDir      = "opt" + string(filepath.Separator) + "proc"          // ProcessDir 进程文件目录
	AppLogDir       = "logs" + string(filepath.Separator) + "app"          // AppLogDir 应用日志目录
)

// HomePath  项目根目录
var HomePath string

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	HomePath = pwd
	if !fileutil.IsExist(ApplicationPath) {
		err := os.MkdirAll(ApplicationPath, os.ModePerm)
		if err != nil {
			global.LOG.Fatal(fmt.Sprintf("创建应用目录失败，错误: %v", err))
		}
	}
	if !fileutil.IsExist(EnvPath) {
		err := os.MkdirAll(EnvPath, os.ModePerm)
		if err != nil {
			global.LOG.Fatal(fmt.Sprintf("创建环境目录失败，错误: %v", err))
		}
	}
	//创建进程目录
	if !fileutil.IsExist(ProcessDir) {
		err := os.MkdirAll(ProcessDir, os.ModePerm)
		if err != nil {
			global.LOG.Fatal(fmt.Sprintf("创建进程目录失败，错误: %v", err))
		}
	}
	if !fileutil.IsExist(AppLogDir) {
		err := os.MkdirAll(AppLogDir, os.ModePerm)
		if err != nil {
			global.LOG.Fatal(fmt.Sprintf("创建应用日志目录失败，错误: %v", err))
		}
	}
}

// GetEnvFullPath 获取环境全路径
func GetEnvFullPath(envDir string) string {
	return filepath.Join(HomePath, EnvPath, envDir)
}

// GetEnvHomePath 获取环境根目录
func GetEnvHomePath() string {
	return filepath.Join(HomePath, EnvPath)
}
func GetLogHomePath() string {
	return filepath.Join(HomePath, "logs")
}

// GetApplicationFullPath 获取应用全路径
func GetApplicationFullPath(appDir string) string {
	return filepath.Join(HomePath, ApplicationPath, appDir)
}

// GetApplicationHomePath 获取应用根目录
func GetApplicationHomePath() string {
	return filepath.Join(HomePath, ApplicationPath)
}

// GetProcessFullPath 获取应用进程全路径
func GetProcessFullPath(appId string) string {
	return filepath.Join(HomePath, ProcessDir, appId)
}

// GetAppLogRelativePath 获取应用相对日志路径
func GetAppLogRelativePath(appId string, appName string) string {
	return filepath.Join(AppLogDir, appName+"_"+appId+".log")
}

// GetAppLogFullPath 获取应用日志全路径
func GetAppLogFullPath(appId string, appName string) string {
	return filepath.Join(HomePath, AppLogDir, appName+"_"+appId+".log")
}

// GetSystemLogFullPath 获取系统日志
func GetSystemLogFullPath() string {
	return filepath.Join(HomePath, LogFilePath)
}
