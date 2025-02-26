package api

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/konglingyinxia/win-start-client/server/dto/res"
	"github.com/konglingyinxia/win-start-client/server/utils/open"
)

type CommManager struct {
}

func NewCommManager() *CommManager {
	return &CommManager{}
}

// OpenFileExplorer 打开文件资源管理器
func (a *CommManager) OpenFileExplorer(dirPath string) res.BaseRes {
	if !fileutil.IsExist(dirPath) {
		return res.Err(fmt.Sprintf("[%s]目录不存在", dirPath))
	}
	err := open.OsFileExplorer(dirPath)
	if err != nil {
		return res.Err(err.Error())
	}
	return res.Ok()
}
