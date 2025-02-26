//go:build windows

package osutils

import (
	"os/exec"
	"syscall"
)

func CmdAddOsParams(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
}
