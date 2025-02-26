package open

import (
	"errors"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/utils/osutils"
	"os/exec"
	"runtime"
)

// OsBrowser 打开默认浏览器
func OsBrowser(url string) {
	var err error
	var cmd *exec.Cmd
	// 判断操作系统并执行相应的命令
	switch runtime.GOOS {
	case "darwin": // macOS
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command(`cmd`, `/c`, `start`, url)
	default:
		global.LOG.Infof("打开默认浏览器失败:不支持的操作系统:%s", runtime.GOOS)
		return
	}
	osutils.CmdAddOsParams(cmd)
	err = cmd.Start()
	if err != nil {
		global.LOG.Errorf("打开默认浏览器失败:%v", err)
	}
}

// OsFileExplorer 打开文件资源管理器
func OsFileExplorer(dirPath string) error {
	var err error
	var cmd *exec.Cmd
	// 判断操作系统并执行相应的命令
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", dirPath)
	case "linux":
		cmd = exec.Command("xdg-open", dirPath)
	case "windows":
		cmd = exec.Command(`cmd`, `/c`, `explorer`, dirPath)
	default:
		global.LOG.Infof("打开文件资源管理器失败:不支持的操作系统:%s", runtime.GOOS)
		return errors.New("不支持的操作系统:" + runtime.GOOS)
	}
	osutils.CmdAddOsParams(cmd)
	err = cmd.Run()
	if err != nil {
		global.LOG.Errorf("打开文件资源管理器失败:%v", err)
	}
	return nil
}
