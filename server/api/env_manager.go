package api

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/dto/req"
	"github.com/konglingyinxia/win-start-client/server/dto/res"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/model"
	"os"
)

type EnvManager struct {
	Separator string `json:"separator"`

	EnvHomeKey  string `json:"envHomeKey"`
	EnvHomePath string `json:"envHomePath"` // 环境根路径

}

func NewEnvManager() *EnvManager {
	return &EnvManager{}
}

// GetEnvHomePath 环境根路径
func (a *EnvManager) GetEnvHomePath() EnvManager {
	return EnvManager{
		Separator:   string(os.PathSeparator),
		EnvHomeKey:  constant.EnvHomeKey,
		EnvHomePath: constant.GetEnvHomePath(),
	}
}

// Add 添加环境
func (a *EnvManager) Add(req req.EnvReqAdd) res.BaseRes {
	if req.EnvName == "" {
		return res.Err("环境名称不能为空")
	}
	env, _ := EnvService.GetOneByName(req.EnvName)
	if env != nil {
		return res.Err("环境名称已存在")
	}
	envDir := constant.GetEnvFullPath(req.EnvDir)
	if !fileutil.IsExist(envDir) {
		return res.Err("环境目录不存在:" + envDir)
	}
	err := EnvService.AddEnv(req)
	if err != nil {
		return res.Err("添加环境失败:" + err.Error())
	}
	return res.Ok()
}

// Delete 删除环境
func (a *EnvManager) Delete(id string) res.BaseRes {
	//查询是否有应用在使用该环境
	count, err := AppService.CountByEnvID(id)
	if err != nil {
		return res.Err("查询应用数量失败:" + err.Error())
	}
	if count > 0 {
		return res.Err("该环境已被应用使用，不能删除")
	}
	err = EnvService.DeleteEnv(id)
	if err != nil {
		return res.Err("删除环境失败:" + err.Error())
	}
	return res.Ok()

}

// List 获取环境列表
func (a *EnvManager) List() res.BaseRes {
	envs, err := EnvService.List()
	if err != nil {
		return res.Err("获取环境列表失败:" + err.Error())
	}
	return res.Success(envs)
}

// Detail 查询环境详情
func (a *EnvManager) Detail(id string) res.BaseRes {
	env, err := EnvService.GetById(id)
	if err != nil {
		return res.Err("查询环境详情失败:" + err.Error())
	}
	r := res.EnvRes{
		BaseModel: model.BaseModel{
			ID:        env.ID,
			CreatedAt: env.CreatedAt,
			UpdatedAt: env.UpdatedAt,
		},
		EnvName: env.EnvName,
		EnvDir:  env.EnvDir,
		EnvVars: env.EnvVars,
		Remark:  env.Remark,
		RootDir: constant.GetEnvFullPath(env.EnvDir),
	}
	return res.Success(r)

}

// Update 更新环境
func (a *EnvManager) Update(req req.EnvReqAdd) res.BaseRes {
	if req.EnvName == "" {
		return res.Err("环境名称不能为空")
	}
	err := EnvService.EditEnv(req)
	if err != nil {
		return res.Err("更新环境失败:" + err.Error())
	}
	return res.Ok()
}

// CheckEnvVars 验证环境变量是否正确
func (a *EnvManager) CheckEnvVars(req req.EnvCheckReq) res.BaseRes {
	env, err := EnvService.GetById(req.ID)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("查询环境详情失败:%s", err.Error()))
		return res.Err("查询环境详情失败:" + err.Error())
	}
	result, err := EnvService.CheckEnvVars(env, req.Command)
	if err != nil {
		global.LOG.Error(fmt.Sprintf("验证环境变量失败:%s", err.Error()))
		result = result + "\n错误结果" + err.Error()
	}
	return res.Success(result)
}
