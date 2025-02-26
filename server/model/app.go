package model

type App struct {
	BaseModel
	Type       string `json:"type"`       // 应用类型:mysql,redis,nginx,custom
	Name       string `json:"name"`       // 应用名称
	Version    string `json:"version"`    // 版本信息
	AppDir     string `json:"appDir"`     // 应用路径(相对根路径)
	EnvId      string `json:"envId"`      // 环境ID
	StartCmd   string `json:"startCmd"`   // 启动命令[]string
	StopCmd    string `json:"stopCmd"`    // 停止命令[]string
	RestartCmd string `json:"restartCmd"` // 重启命令[]string
	VersionCmd string `json:"versionCmd"` // 版本命令[]string
	LogDir     string `json:"logDir"`     // 日志路径
	AutoStart  *bool  `json:"autoStart"`  // 是否自动启动
	StartDelay int    `json:"startDelay"` // 自启启动延迟时间 (单位:秒)
	StartOrder int    `json:"startOrder"` // 自启启动顺序
	Ports      string `json:"ports"`      // 程序占用端口
	EnvVars    string `json:"envVars"`    // 环境变量
	Username   string `json:"username"`
	Password   string `json:"password"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	InitDb     bool   `json:"initDb"` // 是否已经初始化数据库
	Remark     string `json:"remark"` // 备注

}
