package service

import (
	"errors"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/model"
	"gorm.io/gorm"
)

type GlobalSettingService struct{}

type IGlobalSettingService interface {
	GetGlobalSetting() (*model.GlobalSetting, error)
	SetGlobalSetting(req model.GlobalSetting) error
}

func NewGlobalSettingService() IGlobalSettingService {
	return &GlobalSettingService{}
}

func (g GlobalSettingService) GetGlobalSetting() (*model.GlobalSetting, error) {
	var setting *model.GlobalSetting
	err := global.DB.First(&setting).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return setting, err
}

func (g GlobalSettingService) SetGlobalSetting(req model.GlobalSetting) error {
	setting, err := g.GetGlobalSetting()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			req.ID = "0"
			result := global.DB.Create(&req)
			if result.Error != nil {
				return err
			}
			setting = &req
		} else {
			return err
		}
	}
	err = global.DB.Model(&model.GlobalSetting{BaseModel: model.BaseModel{ID: setting.ID}}).Updates(req).Error
	return err
}
