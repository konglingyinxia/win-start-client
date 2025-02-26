package req

type EnvReqAdd struct {
	ID      string `json:"id"`
	EnvName string `json:"envName" binding:"required"`
	EnvDir  string `json:"envDir" binding:"required"`
	EnvVars string `json:"envVars" binding:"required"` // 环境变量
	Remark  string `json:"remark" binding:"required"`  // 备注
}

// EnvCheckReq 验证环境变量
type EnvCheckReq struct {
	ID      string `json:"id" binding:"required"`      // 环境ID
	Command string `json:"command" binding:"required"` // 要执行的命令
}
