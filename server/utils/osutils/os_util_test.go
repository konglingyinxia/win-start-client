package osutils

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestDiskCount(t *testing.T) {
	count := DirSize("/home/data/")
	fmt.Println(count)

}

func TestProcess(t *testing.T) {
	// 启动外部程序
	cmd := exec.Command("sleep", "100") // 这里以sleep命令为例

	// 启动命令并将其标准输出和标准错误输出连接到当前进程
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("启动程序失败:", err)
		return
	}
	fmt.Println("程序已启动，PID:", cmd.Process.Pid)

	// 等待一段时间后停止程序
	time.Sleep(5 * time.Second)

	// 停止外部程序
	if err := cmd.Process.Kill(); err != nil {
		fmt.Println("停止程序失败:", err)
		return
	}

	fmt.Println("程序已停止")
}
