package service

import (
	"bufio"
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/model"
	"github.com/saintfish/chardet"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"time"
)

const ()

type LogService struct {
	appLogReadGo    map[string]bool
	systemLogReadGo map[string]bool
}

type ILogService interface {
	AppLog(app *model.App)
	CloseAppLog(id string)
	SystemLog()
	CloseSystemLog()
}

func NewLogService() ILogService {
	return &LogService{
		appLogReadGo: make(map[string]bool),
		//TODO 系统日志
		systemLogReadGo: make(map[string]bool),
	}
}

func (l LogService) AppLog(app *model.App) {
	//获取应用日志文件
	logFile := constant.GetAppLogFullPath(app.ID, app.Name)
	if fileutil.IsExist(logFile) {
		go loopReadLog(l, logFile, app)
	}
}
func (l LogService) CloseAppLog(id string) {
	l.appLogReadGo[id] = false
}

func loopReadLog(l LogService, logFile string, app *model.App) {
	charset := detectFileCharset(logFile)
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// 设置一个偏移量，表示我们从文件的末尾开始读取
	_, err = file.Seek(-4*1024, io.SeekEnd)
	if err != nil {
		global.LOG.Info(fmt.Sprintf("查找应用【%s】日志文件末尾时出错:%v", app.Name, err))
	}
	if l.appLogReadGo[app.ID] == true {
		global.LOG.Info(fmt.Sprintf("应用【%s】日志监听已开启", app.Name))
		return
	}
	l.appLogReadGo[app.ID] = true
	// 从当前偏移位置读取最后一行
	reader := bufio.NewReader(file)
	for {
		if !l.appLogReadGo[app.ID] {
			global.LOG.Info(fmt.Sprintf("应用【%s】日志监听关闭", app.Name))
			break
		}
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				// 文件结束，等待一段时间再继续尝试
				time.Sleep(300 * time.Millisecond)
				continue
			}
			global.LOG.Error(fmt.Sprintf("应用【%s】日志读取错误: %s", app.Name, err))
			break
		}
		logLine := string(line)
		if charset != "UTF-8" {
			output, _ := simplifiedchinese.GBK.NewDecoder().String(logLine) //转换字符集，解决UTF-8乱码
			if output != "" {
				logLine = output
			}
		}
		runtime.EventsEmit(*global.WailsContext, app.ID, logLine)
	}
}

func (l LogService) SystemLog() {
	//获取应用日志文件
	logFile := constant.GetSystemLogFullPath()
	if fileutil.IsExist(logFile) {
		go loopReadSystemLog(l, logFile)
	}
}
func (l LogService) CloseSystemLog() {
	l.systemLogReadGo[constant.EventSystemLogName] = false
}
func loopReadSystemLog(l LogService, logFile string) {
	file, err := os.Open(logFile)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("打开日志文件出错: %s", err))
		return
	}
	defer file.Close()
	// 设置一个偏移量，表示我们从文件的末尾开始读取
	_, err = file.Seek(-4*1024, io.SeekEnd)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("查找文件末尾时出错: %s", err))
	}
	if l.systemLogReadGo[constant.EventSystemLogName] == true {
		global.LOG.Info("系统日志监听已开启.....")
		return
	}
	global.LOG.Info("系统日志监听开启.....")
	l.systemLogReadGo[constant.EventSystemLogName] = true
	// 从当前偏移位置读取最后一行
	reader := bufio.NewReader(file)
	for {
		if !l.systemLogReadGo[constant.EventSystemLogName] {
			global.LOG.Info("系统日志监听关闭....")
			break
		}
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				// 文件结束，等待一段时间再继续尝试
				time.Sleep(300 * time.Millisecond)
				continue
			}
			global.LOG.Error(fmt.Sprintf("系统日志读取错误: %s", err))
			break
		}
		runtime.EventsEmit(*global.WailsContext, "system-log", string(line))
	}
}
func detectFileCharset(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	buf := make([]byte, 1024)
	_, err = reader.Read(buf)
	if err != nil {
		return ""
	}
	detector := chardet.NewTextDetector()
	charset, err := detector.DetectBest(buf)
	return charset.Charset
}
