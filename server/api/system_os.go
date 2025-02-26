package api

import (
	"errors"
	"fmt"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/dto"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/utils/ip"
	"github.com/konglingyinxia/win-start-client/server/utils/osutils"
	"github.com/shirou/gopsutil/v4/host"
)

type SystemOs struct {
}

func NewSystemOs() *SystemOs {
	return &SystemOs{}
}

// OsInfo  系统信息
func (a *SystemOs) OsInfo() (*dto.SystemInfo, error) {
	info, err := host.Info()
	hostInfo, err := osutils.Host()
	if err != nil {
		global.LOG.Error("获取主机network信息失败：" + err.Error())
		return nil, errors.New("获取主机network信息失败：" + err.Error())
	}
	cpuInfo := hostInfo.CPU
	gpuInfo := hostInfo.GPU
	return &dto.SystemInfo{
		InfoStat:     *info,
		Cpu:          fmt.Sprintf("%s %d(%d)", cpuInfo.Processors[0].Model, cpuInfo.TotalCores, cpuInfo.TotalThreads),
		Gpu:          osutils.GpuParse(gpuInfo),
		Memory:       osutils.MemParse(),
		Network:      hostInfo.Network.String(),
		Chassis:      hostInfo.Chassis.String(),
		Product:      hostInfo.Product.String(),
		Baseboard:    hostInfo.Baseboard.String(),
		Bios:         hostInfo.BIOS.String(),
		Disk:         osutils.DiskParse(),
		Manufacturer: hostInfo.Baseboard.Vendor,
		ProductName:  hostInfo.Product.Name,
		RootDir:      constant.HomePath,
		IpAddr:       ip.LocalOutboundIp(),
	}, nil
}

// HomeOsInfo 获取首页系统信息
func (a *SystemOs) HomeOsInfo() (*dto.HomeOsInfo, error) {
	info, err := host.Info()
	if err != nil {
		global.LOG.Error("获取主机network信息失败：" + err.Error())
		return nil, errors.New("获取主机network信息失败：" + err.Error())
	}
	return &dto.HomeOsInfo{
		InfoStat: *info,
		IpAddr:   ip.LocalOutboundIp(),
	}, nil
}

// Memory 内存占用
func (a *SystemOs) Memory() string {
	return "hello world"
}

// Disk 磁盘信息
func (a *SystemOs) Disk() string {
	return "hello world"
}

// Cpu cpu信息
func (a *SystemOs) Cpu() string {
	return "hello world"
}

// Gpu gpu信息
func (a *SystemOs) Gpu() string {
	return "hello world"
}

// Network 网络信息
func (a *SystemOs) Network() string {
	return "hello world"
}

// Process 进程信息
func (a *SystemOs) Process() string {
	return "hello world"
}
