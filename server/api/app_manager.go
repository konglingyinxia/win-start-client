package api

import (
	"encoding/json"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/dto/req"
	"github.com/konglingyinxia/win-start-client/server/dto/res"
	"github.com/konglingyinxia/win-start-client/server/model"
	"os"
)

type AppManager struct {
	Separator   string `json:"separator"`
	AppHomeKey  string `json:"appHomeKey"`
	AppHomePath string `json:"appHomePath"`
}

func NewAppManager() *AppManager {
	return &AppManager{}
}

// GetAppHomePath 环境根路径
func (a *AppManager) GetAppHomePath() AppManager {
	return AppManager{
		Separator:   string(os.PathSeparator),
		AppHomeKey:  constant.AppHomeKey,
		AppHomePath: constant.GetApplicationHomePath(),
	}
}

// Add 添加
func (a *AppManager) Add(req req.AppReq) res.BaseRes {
	if req.AppDir == "" {
		return res.Err("appDir不能为空")
	}
	if req.Type == "" {
		return res.Err("type不能为空")
	}
	if req.Name == "" {
		return res.Err("name不能为空")
	}
	rootDir := constant.GetApplicationFullPath(req.AppDir)
	if !fileutil.IsDir(rootDir) {
		return res.Err("appDir不存在")
	}
	if req.Type == constant.AppTypeMysql && req.Port == 0 {
		return res.Err("port不能为空")
	}
	err := AppService.Add(req)
	if err != nil {
		return res.Err(err.Error())
	}
	return res.Ok()
}

// List 列表
func (a *AppManager) List() res.BaseRes {
	apps, err := AppService.List()
	if err != nil {
		return res.Err("获取环境列表失败:" + err.Error())
	}
	return res.Success(apps)
}

// GetById 查询一个
func (a *AppManager) GetById(id string) res.BaseRes {
	app, err := AppService.GetById(id)
	if err != nil {
		return res.Err("获取环境失败:" + err.Error())
	}
	return res.Success(app)
}

// Delete 删除
func (a *AppManager) Delete(id string) res.BaseRes {
	err := AppService.Delete(id)
	if err != nil {
		return res.Err("删除环境失败:" + err.Error())
	}
	return res.Ok()
}

// Update 更新
func (a *AppManager) Update(req req.AppReq) res.BaseRes {
	if req.ID == "" {
		return res.Err("id不能为空")
	}
	err := AppService.Update(req)
	if err != nil {
		return res.Err("更新环境失败:" + err.Error())
	}
	return res.Ok()

}

// ExecuteSQL 执行sql命令
func (a *AppManager) ExecuteSQL() res.BaseRes {
	return res.Ok()

}

//===================控制指令==============

// Start 启动
func (a *AppManager) Start(id string) res.BaseRes {
	app, err := AppService.GetById(id)
	if err != nil {
		return res.Err("获取应用失败:" + err.Error())
	}
	err = AppService.Start(app)
	if err != nil {
		return res.Err("启动失败:" + err.Error())
	}
	return res.Ok()
}

// StartAll 全部启动
func (a *AppManager) StartAll(ids []string) res.BaseRes {
	apps, err := AppService.ListByIds(ids)
	if err != nil {
		return res.Err("获取应用失败:" + err.Error())
	}
	err = AppService.StartAll(false, apps)
	if err != nil {
		return res.Err("应用启动失败:" + err.Error())
	}
	return res.Ok()

}

// Stop 暂停
func (a *AppManager) Stop(id string) res.BaseRes {
	app, err := AppService.GetById(id)
	if err != nil {
		return res.Err("获取应用失败:" + err.Error())
	}
	err = AppService.Stop(app)
	if err != nil {
		return res.Err("暂停失败:" + err.Error())
	}
	return res.Ok()
}

// StopAll 全部暂停
func (a *AppManager) StopAll(ids []string) res.BaseRes {
	apps, err := AppService.ListByIds(ids)
	if err != nil {
		return res.Err("获取应用失败:" + err.Error())
	}
	err = AppService.StopAll(false, apps)
	if err != nil {
		return res.Err("应用启动失败:" + err.Error())
	}
	return res.Ok()
}

// Restart 重启
func (a *AppManager) Restart(id string) res.BaseRes {
	app, err := AppService.GetById(id)
	if err != nil {
		return res.Err("获取应用失败:" + err.Error())
	}
	err = AppService.Restart(app)
	if err != nil {
		return res.Err("暂停失败:" + err.Error())
	}
	return res.Ok()
}

// RestartAll 全部重启
func (a *AppManager) RestartAll() res.BaseRes {
	return res.Ok()
}

// =============查看信息=================

// Process 进程信息
func (a *AppManager) Process() res.BaseRes {
	return res.Ok()

}

// Status 状态
func (a *AppManager) Status(id string) res.BaseRes {
	app, err := AppService.GetById(id)
	if err != nil {
		return res.Err("获取应用失败:" + err.Error())
	}
	status, err := AppService.RunningStatus([]model.App{*app})
	if err != nil {
		return res.Err("获取状态失败:" + err.Error())
	}
	if len(status) == 0 {
		return res.Success(res.AppRunStatus{})
	}
	return res.Success(status[0])
}

// RefreshPorts 刷新端口
func (a *AppManager) RefreshPorts(id string) res.BaseRes {
	app, err := AppService.GetById(id)
	if err != nil {
		return res.Err("获取应用失败:" + err.Error())
	}
	ports, err := AppService.RefreshPorts(*app)
	if err != nil {
		return res.Err("刷新程序端口失败:" + err.Error())
	}
	marshal, err := json.Marshal(ports)
	if err != nil {
		return res.Err("序列化失败:" + err.Error())
	}
	return res.Success(string(marshal))
}
