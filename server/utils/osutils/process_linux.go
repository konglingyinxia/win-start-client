// SPDX-License-Identifier: BSD-3-Clause
//go:build linux

package osutils

import (
	"errors"
	"github.com/shirou/gopsutil/v4/process"
)

func KillProcess(pids []uint32) error {
	return errors.New("not implemented yet")
}
func ProcGroupPids(parentPid uint32) ([]uint32, error) {
	return []uint32{}, errors.New("not implemented yet")
}

func ProcChildren(param *process.Process) ([]*process.Process, error) {
	return param.Children()
}
