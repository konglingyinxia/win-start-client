package api

import (
	"errors"
	"fmt"
	"github.com/konglingyinxia/win-start-client/server/dto/res"
	"github.com/konglingyinxia/win-start-client/server/model"
	"gorm.io/gorm"
)

type GlobalSettingManager struct{}

func NewGlobalSettingManager() *GlobalSettingManager {
	return &GlobalSettingManager{}
}

// GetGlobalSetting 获取全局设置
func (gs *GlobalSettingManager) GetGlobalSetting() res.BaseRes {
	result, err := GlobalSettingService.GetGlobalSetting()
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return res.Err(fmt.Sprintf("获取全局设置失败:%s", err))
		}
	}
	return res.Success(result)
}
func (gs *GlobalSettingManager) SetGlobalSetting(req model.GlobalSetting) res.BaseRes {
	err := GlobalSettingService.SetGlobalSetting(req)
	if err != nil {
		return res.Err(fmt.Sprintf("设置全局设置失败:%s", err))
	}
	return res.Ok()
}
