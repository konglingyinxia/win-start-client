//go:build windows

package singleton

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/ncruces/zenity"
	"golang.org/x/sys/windows"
	"os"
	"syscall"
)

func CheckRunning() (running bool) {
	path, err := os.Executable()
	if err != nil {
		global.LOG.Error("获取程序路径失败,程序退出！", err)
		zenity.Warning("获取程序路径失败,程序退出！", zenity.ErrorIcon, zenity.Title(constant.WinTitle))
		os.Exit(1)
		return false
	}
	hashName := md5.Sum([]byte(path))
	name, err := syscall.UTF16PtrFromString("Global\\" + hex.EncodeToString(hashName[:]))
	if err != nil {
		global.LOG.Error("检测多开进程失败,程序退出", err)
		zenity.Warning("检测多开进程失败,程序退出", zenity.ErrorIcon, zenity.Title(constant.WinTitle))
		os.Exit(1)
		return false
	}
	mutex, err := windows.CreateMutex(nil, false, name)
	if err != nil {
		global.LOG.Error("获取多开进程状态失败", err)
		return true
	}
	return mutex <= 0

}

func ReleaseLock() {

}

func AlertWaring(text string) {
	err := zenity.Warning(text,
		zenity.Title(constant.WinTitle),
		zenity.InfoIcon)
	if err != nil {
		global.LOG.Error("弹窗提示：程序运行中失败", err)
		return
	}
}
