package api

import "github.com/konglingyinxia/win-start-client/server/service"

var (
	EnvService           = service.NewEnvService()
	AppService           = service.NewAppService()
	LogService           = service.NewLogService()
	GlobalSettingService = service.NewGlobalSettingService()
)
