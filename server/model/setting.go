package model

// GlobalSetting 全局设置
type GlobalSetting struct {
	BaseModel
	DefaultWeb     string `json:"defaultWeb"`     //默认网页
	OpenDefaultWeb *bool  `json:"openDefaultWeb"` //是否打开默认网页
}
