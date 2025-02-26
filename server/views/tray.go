package views

import (
	"github.com/duke-git/lancet/v2/system"
	"github.com/getlantern/systray"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/global"
	runtime2 "github.com/wailsapp/wails/v2/pkg/runtime"
)

// RegisterTray 启动托盘
func RegisterTray() {
	systray.Register(onReadyTray, onExitTray)
}
func StartTray() {
	systray.Run(onReadyTray, onExitTray)
}
func onReadyTray() {
	if system.IsWindows() {
		// 托盘图标
		systray.SetIcon(global.Icon)
	} else {
		// 托盘图标
		systray.SetIcon(global.PNGIcon)
	}
	systray.SetTitle(constant.WinTitle)
	systray.SetTooltip(constant.WinTitle)
	// 托盘菜单
	onTrayMenuTray()
}
func onTrayMenuTray() {
	// 创建菜单项
	mShow := systray.AddMenuItem("显示", "显示应用")
	mHide := systray.AddMenuItem("隐藏", "隐藏应用")
	mQuit := systray.AddMenuItem("退出", "退出应用")
	// 处理菜单项的点击事件
	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				global.LOG.Info("显示应用")
				runtime2.Show(*global.WailsContext)
				runtime2.WindowCenter(*global.WailsContext)
			case <-mHide.ClickedCh:
				global.LOG.Info("隐藏应用")
				runtime2.Hide(*global.WailsContext)
			case <-mQuit.ClickedCh:
				global.LOG.Info("退出应用")
				runtime2.Show(*global.WailsContext)
				runtime2.WindowCenter(*global.WailsContext)
				runtime2.Quit(*global.WailsContext)
			}
		}
	}()
}

func onExitTray() {
}
