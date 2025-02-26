package res

import "github.com/konglingyinxia/win-start-client/server/model"

type EnvRes struct {
	model.BaseModel
	RootDir string `json:"rootDir"` // 根目录
	EnvName string `json:"envName"` // 环境名称
	EnvDir  string `json:"envDir"`  // 环境路径(相对根路径)
	EnvVars string `json:"envVars"` // 环境变量
	Remark  string `json:"remark"`  // 备注
}
