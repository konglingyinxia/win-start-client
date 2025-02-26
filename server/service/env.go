package service

import (
	"encoding/json"
	"errors"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/dto/req"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/model"
	env2 "github.com/konglingyinxia/win-start-client/server/utils/env"
	"github.com/konglingyinxia/win-start-client/server/utils/id"
	"gorm.io/gorm"
	"os"
	"os/exec"
	"strings"
)

type EnvService struct{}

type IEnvService interface {
	AddEnv(req req.EnvReqAdd) error
	// EditEnv 编辑环境变量
	EditEnv(req req.EnvReqAdd) error
	// DeleteEnv 删除环境变量
	DeleteEnv(id string) error
	GetOneByName(name string) (*model.Env, error)
	List() (*[]model.Env, error)
	GetById(id string) (*model.Env, error)
	CheckEnvVars(env *model.Env, command string) (string, error)
	CountBy() int64
}

func NewEnvService() IEnvService {
	return &EnvService{}
}

func (e EnvService) List() (*[]model.Env, error) {
	var envs []model.Env
	err := global.DB.
		Select(
			"id", "env_name", "env_dir", "remark", "created_at", "updated_at",
		).Find(&envs).Error
	return &envs, err
}
func (e EnvService) AddEnv(req req.EnvReqAdd) error {
	envVars := "{}"
	if req.EnvVars != "" {
		err := json.Unmarshal([]byte(req.EnvVars), &map[string]string{})
		if err != nil {
			return errors.New("环境变量格式错误:" + err.Error())
		}
		envVars = req.EnvVars
	}
	env := model.Env{
		EnvName: req.EnvName,
		EnvDir:  req.EnvDir,
		Remark:  req.Remark,
		EnvVars: envVars,
		BaseModel: model.BaseModel{
			ID: id.UuidSimple(),
		},
	}
	err := global.DB.Create(&env).Error
	if err != nil {
		return err
	}
	return nil
}

func (e EnvService) EditEnv(req req.EnvReqAdd) error {
	env, err := e.GetById(req.ID)
	if err != nil {
		return err
	}
	if env != nil && env.ID != req.ID {
		return errors.New("环境变量名称已存在")
	}
	env.ID = req.ID
	env.EnvName = req.EnvName
	env.EnvDir = req.EnvDir
	env.Remark = req.Remark
	env.EnvVars = req.EnvVars
	err = global.DB.Model(&model.Env{}).Where("id = ?", req.ID).Updates(env).Error
	return err

}

func (e EnvService) DeleteEnv(id string) error {
	err := global.DB.Delete(&model.Env{BaseModel: model.BaseModel{ID: id}}).Error
	return err
}

func (e EnvService) GetOneByName(name string) (*model.Env, error) {
	env := model.Env{}
	err := global.DB.Where("env_name = ?", name).First(&env).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return &env, err
}

func (e EnvService) GetById(id string) (*model.Env, error) {
	var env model.Env
	err := global.DB.First(&env, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return &env, err
}

func (e EnvService) CheckEnvVars(env *model.Env, command string) (string, error) {
	//空格分割Command
	commands := strings.Split(command, " ")
	//内置环境变量
	var envMaps map[string]string
	if env.EnvVars != "" {
		err := json.Unmarshal([]byte(env.EnvVars), &envMaps)
		if err != nil {
			return "", errors.New("环境变量格式错误:" + err.Error())
		}
	}
	//解析环境变量/
	resultEnv := env2.ParseAllEnv(envMaps)
	for key, val := range resultEnv {
		os.Setenv(key, val)
	}
	//执行命令
	var cmd *exec.Cmd
	if len(commands) == 1 {
		cmd = exec.Command(commands[0])
	} else {
		cmd = exec.Command(commands[0], commands[1:]...)
	}
	//设置执行目录
	cmd.Dir = constant.GetEnvFullPath(env.EnvDir)
	//获取命令输出
	out, err := cmd.CombinedOutput()
	return string(out), err
}
func (e EnvService) CountBy() int64 {
	var count int64
	global.DB.Find(&model.Env{}).Count(&count)
	return count
}
