package views

import (
	"context"
	"github.com/duke-git/lancet/v2/system"
	"github.com/getlantern/systray"
	"github.com/konglingyinxia/win-start-client/server/api"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/dto/res"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/model"
	"github.com/konglingyinxia/win-start-client/server/utils/open"
	runtime2 "github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

//运行管理

// RuntimeApp struct
type RuntimeApp struct {
	ctx context.Context
}

// NewRuntimeApp  creates a new App application struct
func NewRuntimeApp() *RuntimeApp {
	return &RuntimeApp{}
}

// Startup 在应用程序启动时调用。上下文已保存
// so we can call the runtime methods
func (a *RuntimeApp) Startup(ctx context.Context) {
	a.ctx = ctx
	global.WailsContext = &ctx
	global.LOG.Info("注册系统托盘")
	if system.IsWindows() {
		go StartTray()
	} else {
		RegisterTray()
	}
	global.LOG.Info("注册系统托盘-完成")
}

// OnDomReady 在DOM准备就绪时调用。
func (a *RuntimeApp) OnDomReady(ctx context.Context) {
	time.Sleep(1 * time.Second)
	runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName,
		res.AppStartEvent{
			EventStatus: "running",
			EventType:   "start",
		})
	//执行启动所有应用服务
	err := api.AppService.StartAll(true, nil)
	if err != nil {
		global.LOG.Error("启动应用服务失败:%s", err)
		runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName,
			res.AppStartEvent{
				EventStatus: "end",
				EventType:   "start",
				Error:       err.Error(),
			})
	} else {
		runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName,
			res.AppStartEvent{
				EventStatus: "end",
				EventType:   "start",
			})
		//打开默认浏览器
		setting, _ := api.GlobalSettingService.GetGlobalSetting()
		if setting == nil {
			flag := false
			setting = &model.GlobalSetting{
				OpenDefaultWeb: &flag,
				DefaultWeb:     "http://127.0.0.1:8005",
			}
		}
		//打开默认浏览器
		if setting.OpenDefaultWeb != nil && *setting.OpenDefaultWeb {
			openUrl := "http://127.0.0.1:8005"
			if setting.DefaultWeb != "" {
				openUrl = setting.DefaultWeb
			}
			global.LOG.Info("打开默认浏览器:", openUrl)
			open.OsBrowser(openUrl)
		}
	}
}
func (a *RuntimeApp) OnBeforeClose(ctx context.Context) (prevent bool) {
	dialog, err := runtime2.MessageDialog(ctx, runtime2.MessageDialogOptions{
		Type:    runtime2.QuestionDialog,
		Title:   "退出",
		Message: "您确定要退出吗?",
	})
	if err != nil {
		runtime2.MessageDialog(ctx, runtime2.MessageDialogOptions{
			Type:    runtime2.ErrorDialog,
			Title:   "提示",
			Message: "退出失败:" + err.Error(),
		})
		return true
	}
	if dialog == "Yes" {
		runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName,
			res.AppStartEvent{
				EventStatus: "running",
				EventType:   "stop",
			})
		//执行停止所有应用服务
		err = api.AppService.StopAll(true, nil)
		if err != nil {
			global.LOG.Error("停止应用服务失败:%s", err)
			runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName,
				res.AppStartEvent{
					EventStatus: "end",
					EventType:   "stop",
					Error:       err.Error(),
				})
		} else {
			runtime2.EventsEmit(*global.WailsContext, constant.EventAppStartStopName,
				res.AppStartEvent{
					EventStatus: "end",
					EventType:   "stop",
				})
		}
		time.Sleep(2 * time.Second)
		systray.Quit()
		global.LOG.Info("退出程序")
		return false
	}
	return true
}

// Shutdown 在应用程序结束时调用。
func (a *RuntimeApp) Shutdown(ctx context.Context) {

}
