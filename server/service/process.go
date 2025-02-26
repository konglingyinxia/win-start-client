package service

import (
	"github.com/duke-git/lancet/v2/slice"
	"github.com/shirou/gopsutil/v4/process"
)

type iProcessService struct{}

var ProcessService iProcessService

func init() {
	ProcessService = iProcessService{}
}

func (processService *iProcessService) CreateProcess() (err error) {
	return
}

// ProcessPorts  获取进程占用的端口
func (processService *iProcessService) ProcessPorts(proc *process.Process) (ports []int) {
	connections, err := proc.Connections()
	if err != nil {
		return []int{}
	}
	for _, value := range connections {
		if value.Status == "LISTEN" {
			return append(ports, int(value.Laddr.Port))
		}
	}
	return
}

// ProcessPortsByPid  获取进程占用的端口
func (processService *iProcessService) ProcessPortsByPid(pid int32) (ports []int) {
	proc, err := process.NewProcess(pid)
	if err != nil {
		return []int{}
	}
	connections, err := proc.Connections()
	if err != nil {
		return []int{}
	}
	for _, value := range connections {
		if value.Status == "LISTEN" {
			port := int(value.Laddr.Port)
			if !slice.Contain(ports, port) {
				ports = append(ports, int(value.Laddr.Port))
			}
		}
	}
	return
}
