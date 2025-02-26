package main

import (
	"embed"
	"fmt"
	"github.com/konglingyinxia/win-start-client/server"
	"github.com/konglingyinxia/win-start-client/server/api"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/utils"
	"github.com/konglingyinxia/win-start-client/server/utils/singleton"
	"github.com/konglingyinxia/win-start-client/server/views"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"os"
	"runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.ico
var icon []byte

//go:embed build/appicon.png
var pngIcon []byte

func init() {
	global.Icon = icon
	global.PNGIcon = pngIcon
	global.AssetsUI = assets
}

var version = "dev"
var buildTime = "unknown"
var buildDate = "unknown"
var buildGoVersion = "unknown"

func main() {
	//1、版本信息
	if len(os.Args) > 1 && (os.Args[1] == "--version" ||
		os.Args[1] == "-v") {
		buildInfo()
		return
	}
	buildInfo()
	//2、数据初始化
	server.Start()
	//3、单进程判断
	if singleton.CheckRunning() {
		alertMsg := "程序已在运行中，请勿重复启动！"
		singleton.AlertWaring(alertMsg)
		global.LOG.Info(alertMsg)
		return
	}
	defer singleton.ReleaseLock()
	//4、启动窗口
	startWindow()
}

// startWindow 启动窗口
func startWindow() {
	app := views.NewRuntimeApp()
	err := wails.Run(&options.App{
		Title:  constant.WinTitle,
		Width:  constant.WinWidth,
		Height: constant.WinHeight,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Logger:            logger.NewFileLogger(constant.LogFilePath),
		LogLevel:          logger.INFO,
		HideWindowOnClose: true,
		OnStartup:         app.Startup,
		OnDomReady:        app.OnDomReady,
		OnShutdown:        app.Shutdown,
		OnBeforeClose:     app.OnBeforeClose,
		Bind: []interface{}{
			api.NewSystemOs(),
			api.NewAppManager(),
			api.NewEnvManager(),
			api.NewLogManager(),
			api.NewGlobalSettingManager(),
			api.NewDashboardManager(),
			api.NewCommManager(),
		},
		Linux: &linux.Options{
			Icon: icon,
		},
		Windows: &windows.Options{
			WebviewBrowserPath: utils.WebviewQuery(),
		},
	})
	if err != nil {
		global.LOG.Error("窗口启动失败:" + err.Error())
	} else {
		global.LOG.Info("窗口启动成功...")
	}
}

func buildInfo() {
	fmt.Println("Server Version      :", version)
	fmt.Println("Build  Time         :", buildDate+" "+buildTime)
	fmt.Println("Build Go Version    :", buildGoVersion)
	fmt.Println("GOARCH              :", runtime.GOARCH)
	fmt.Println("GOOS                :", runtime.GOOS)
	fmt.Println("NumCPU              :", runtime.NumCPU())
	fmt.Println("GO_ROOT             :", runtime.GOROOT())

}
