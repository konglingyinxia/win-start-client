package res

import (
	"github.com/konglingyinxia/win-start-client/server/model"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/process"
)

type AppRes struct {
	model.BaseModel
	Type         string       `json:"type"`                      // 应用类型:mysql,redis,nginx,custom
	AutoStart    bool         `json:"autoStart"`                 // 是否自动启动
	StartDelay   int          `json:"startDelay"`                // 自启启动延迟时间 (单位:秒)
	StartOrder   int          `json:"startOrder"`                // 自启启动顺序
	Ports        string       `json:"ports"`                     // 程序占用端口
	Name         string       `json:"name"`                      // 应用名称
	Version      string       `json:"version"`                   // 应用版本
	RootDir      string       `json:"rootDir"`                   // 根目录
	AppDir       string       `json:"appDir"`                    // 应用路径
	LogDir       string       `json:"logDir"`                    // 日志路径
	DiskSize     int64        `json:"diskSize"`                  // 应用磁盘占用
	Remark       string       `json:"remark" binding:"required"` // 备注
	AppRunStatus AppRunStatus `json:"runStatus"`                 // 应用运行状态
}
type AppRunStatus struct {
	Status     string                  `json:"status"`     // 应用状态: running, padding, stopped
	Pid        int                     `json:"pid"`        // 应用进程ID
	Pids       []int                   `json:"pids"`       // 应用进程ID
	Memory     *process.MemoryInfoStat `json:"memory"`     // 应用内存占用
	MemSize    uint64                  `json:"memSize"`    // 应用内存占用
	MemPercent float32                 `json:"memPercent"` // 应用内存占用百分比
	CpuTime    *cpu.TimesStat          `json:"cpuTime"`    // 应用CPU占用
	CpuPercent float64                 `json:"cpuPercent"` // 应用CPU占用百分比
	Uptime     string                  `json:"uptime"`     // 应用运行时间
}

// AppStartEvent 应用启动/停止状态事件通知
type AppStartEvent struct {
	model.BaseModel
	Status      string `json:"status"`      // 应用状态: started, running, stopping, stopped, failed
	EventType   string `json:"eventType"`   // 事件类型: start, stop
	EventStatus string `json:"eventStatus"` // 事件状态: running-执行中, end-已结束
	Error       string `json:"error"`       // 错误信息
}
