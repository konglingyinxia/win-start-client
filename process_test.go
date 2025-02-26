package main

import (
	"fmt"
	"github.com/shirou/gopsutil/v4/process"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestStartJava(t *testing.T) {
	// 启动外部程序
	cmd := exec.Command("/bin/java", "-jar", "/home/kongling/work/work/project/go/win-start-client/opt/app/adpublish/mzy-model-ai-conference-api-0.0.1-SNAPSHOT.jar") // 这里以sleep命令为例
	cmd.Env = append(os.Environ(), "JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64")

	// 启动命令并将其标准输出和标准错误输出连接到当前进程
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("启动程序失败:", err)
		return
	}
	fmt.Println("程序已启动，PID:", cmd.Process.Pid)

	// 等待一段时间后停止程序
	time.Sleep(10 * time.Second)
	pid := cmd.Process.Pid
	proc, err := process.NewProcess(int32(pid))
	if err != nil {
		return
	}
	connectionsMax, err := proc.Connections()
	if err != nil {
		return
	}
	for _, conn := range connectionsMax {
		fmt.Println(conn)
	}
	fmt.Println("程序连接数:", len(connectionsMax))

	// 停止外部程序
	if err := cmd.Process.Kill(); err != nil {
		fmt.Println("停止程序失败:", err)
		return
	}

	fmt.Println("程序已停止")

}
func TestStart(t *testing.T) {
	var str []string
	fmt.Println(str[0])
	fmt.Println(str[1:])
}
