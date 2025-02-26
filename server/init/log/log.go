package log

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/global"
	log2 "github.com/labstack/gommon/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"path"
)

func Init() {
	logFile := constant.LogFilePath
	if !fileutil.IsExist(logFile) {
		err := fileutil.CreateDir(path.Dir(logFile))
		if err != nil {
			log.Fatal("日志文件创建失败")
			return
		}
		fileutil.CreateFile(logFile)
	}
	var lumLog = NewLumberlog(logFile)
	mul := io.MultiWriter(lumLog, os.Stdout)
	log.SetOutput(mul)
	echoLog := log2.New("-")
	echoLog.SetHeader("${time_rfc3339}  ${level} --- ${short_file} ${line} ${message}")
	echoLog.EnableColor()
	echoLog.SetOutput(mul)
	global.LOG = echoLog
}

func NewLumberlog(fileName string) *lumberjack.Logger {
	var lumLog = &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    5, // 5MB
		MaxBackups: 5,
		MaxAge:     28, // 28 days
		Compress:   true,
	}
	return lumLog
}
