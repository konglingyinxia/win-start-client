package dto

import (
	"github.com/shirou/gopsutil/v4/host"
)

type SystemInfo struct {
	InfoStat     host.InfoStat `json:"infoStat"`     // 系统信息
	IpAddr       string        `json:"ipAddr"`       // ip地址
	RootDir      string        `json:"rootDir"`      //项目路径
	Cpu          string        `json:"cpu"`          // 处理器
	Gpu          string        `json:"gpu"`          // 显卡
	Memory       interface{}   `json:"memory"`       // 内存
	Chassis      string        `json:"chassis"`      // 底盘
	Disk         DiskCount     `json:"disk"`         // 硬盘
	Network      string        `json:"network"`      // 网络
	Product      string        `json:"product"`      // 产品信息
	Baseboard    string        `json:"baseboard"`    // 主板信息
	Bios         string        `json:"bios"`         // bios信息
	Manufacturer string        `json:"manufacturer"` // 厂商信息
	ProductName  string        `json:"productName"`  // 产品名称
}

// DiskCount 磁盘统计信息
type DiskCount struct {
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

type HomeOsInfo struct {
	InfoStat host.InfoStat `json:"infoStat"` // 系统信息
	IpAddr   string        `json:"ipAddr"`   // ip地址
}
