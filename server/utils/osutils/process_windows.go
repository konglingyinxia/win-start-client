// SPDX-License-Identifier: BSD-3-Clause
//go:build windows

package osutils

import (
	"fmt"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/shirou/gopsutil/v4/process"
	"os"
	"syscall"
	"unsafe"
)

var (
	kernel32                     = syscall.NewLazyDLL("kernel32.dll")
	procOpenProcess              = kernel32.NewProc("OpenProcess")
	procTerminateProcess         = kernel32.NewProc("TerminateProcess")
	procCloseHandle              = kernel32.NewProc("CloseHandle")
	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32First           = kernel32.NewProc("Process32First")
	procProcess32Next            = kernel32.NewProc("Process32Next")
)

const (
	ProcessTerminate        = 0x0001
	ProcessQueryInformation = 0x0400
	Th32csSnapprocess       = 0x00000002
)

type PROCESSENTRY32 struct {
	dwSize              uint32
	cntUsage            uint32
	th32ProcessID       uint32
	th32DefaultHeapID   uintptr
	th32ModuleID        uint32
	cntThreads          uint32
	th32ParentProcessID uint32
	pcPriClassBase      int32
	dwFlags             uint32
	szExeFile           [260]uint16
}

func KillProcess(pids []uint32) error {
	for _, pid := range pids {
		_, err := os.FindProcess(int(pid))
		if err != nil {
			continue
		}
		// 打开主进程
		hProcess, _, err := procOpenProcess.Call(
			uintptr(ProcessTerminate|ProcessQueryInformation),
			0,
			uintptr(pid),
		)
		if hProcess == 0 {
			return fmt.Errorf("无法打开进程: %v", err)
		}
		// 终止主进程
		ret, _, err := procTerminateProcess.Call(hProcess, 0)
		if ret == 0 {
			procCloseHandle.Call(hProcess)
			return fmt.Errorf("无法终止进程: %v", err)
		}
		// 关闭进程句柄
		procCloseHandle.Call(hProcess)
	}
	return nil
}

// ProcGroupPids 获取进程组id
func ProcGroupPids(parentPid uint32) ([]uint32, error) {
	var pids []uint32
	// 创建进程快照
	hSnapshot, _, err := procCreateToolhelp32Snapshot.Call(uintptr(Th32csSnapprocess), 0)
	if hSnapshot == 0 {
		return nil, fmt.Errorf("无法创建进程快照: %v", err)
	}
	// 初始化 PROCESSENTRY32 结构体
	pe32 := PROCESSENTRY32{}
	pe32.dwSize = uint32(unsafe.Sizeof(pe32))
	// 获取第一个进程
	success, _, _ := procProcess32First.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
	if success == 0 {
		procCloseHandle.Call(hSnapshot)
		return nil, fmt.Errorf("无法获取第一个进程")
	}
	// 遍历所有进程
	for success != 0 {
		if pe32.th32ParentProcessID == parentPid {
			pids = append(pids, pe32.th32ProcessID)
		}
		success, _, _ = procProcess32Next.Call(hSnapshot, uintptr(unsafe.Pointer(&pe32)))
	}
	// 关闭快照句柄
	procCloseHandle.Call(hSnapshot)
	return pids, nil
}

func ProcChildren(param *process.Process) ([]*process.Process, error) {
	pids, err := ProcGroupPids(uint32(param.Pid))
	if err != nil {
		return nil, fmt.Errorf("无法获取子进程:%v", err)
	}
	var procs []*process.Process
	for _, item := range pids {
		proc, err := process.NewProcess(int32(item))
		if err != nil {
			global.LOG.Errorf("进程信息获取失败：%v", err)
		} else {
			procs = append(procs, proc)
		}
	}
	return procs, nil
}
