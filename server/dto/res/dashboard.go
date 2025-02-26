package res

// OverviewData 概览数据
type OverviewData struct {
	RunningNum int   `json:"runningNum"` //运行中数量
	StoppedNum int   `json:"stoppedNum"` //停止数量
	EnvNum     int64 `json:"envNum"`     //环境数量
	StorageApp int64 `json:"storageApp"` //存储大小
	StorageEnv int64 `json:"storageEnv"` //环境存储大小
	StorageLog int64 `json:"storageLog"` //日志存储大小
}
