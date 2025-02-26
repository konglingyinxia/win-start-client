//go:build linux

package singleton

import (
	"github.com/gofrs/flock"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/ncruces/zenity"
)

// 程序锁定文件
const lockFile = "/tmp/win-start-client.lock"

// 解锁
var fileLock = flock.New(lockFile)

// CheckRunning 检查是否已有程序在运行
func CheckRunning() bool {
	locked, err := fileLock.TryLock()
	if err != nil {
		global.LOG.Error("检查进程锁失败: ", err)
		return false
	}
	//获取锁成功
	if locked {
		//程序未运行
		global.LOG.Info("获取进程锁成功")
		return false
	}
	return true
}

// ReleaseLock 释放锁文件
func ReleaseLock() {
	global.LOG.Info("释放进程锁…")
	err := fileLock.Unlock()
	if err != nil {
		global.LOG.Error("释放进程锁失败: ", err)
		return
	}
}

func AlertWaring(text string) {
	err := zenity.Notify(text,
		zenity.Title(constant.WinTitle),
		zenity.InfoIcon)
	if err != nil {
		global.LOG.Error("通知提示：程序运行中失败", err)
		return
	}
}
