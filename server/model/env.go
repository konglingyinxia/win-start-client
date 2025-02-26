package model

type Env struct {
	BaseModel
	EnvName    string `json:"envName"`    // 环境名称
	EnvDir     string `json:"envDir"`     // 环境路径(相对根路径)
	EnvVars    string `json:"envVars"`    // 环境变量
	Remark     string `json:"remark"`     // 备注
	VersionCmd string `json:"versionCmd"` // 版本命令
}
