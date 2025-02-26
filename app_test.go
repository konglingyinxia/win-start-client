package main

import (
	"bufio"
	"fmt"
	"github.com/konglingyinxia/win-start-client/server"
	"github.com/konglingyinxia/win-start-client/server/api"
	"github.com/konglingyinxia/win-start-client/server/global"
	"github.com/konglingyinxia/win-start-client/server/model"
	"github.com/konglingyinxia/win-start-client/server/utils/osutils"
	"io"
	"log"
	"os"
	"testing"
)

func TestStartApp(t *testing.T) {
	server.Start()
	api.NewAppManager().Start("fb9550a9a9344cf79a27035c07cd7bff")
}
func TestStartStopApp(t *testing.T) {
	server.Start()
	api.NewAppManager().Stop("2a7d6def46b64e059a869abdfe662990")
}
func TestRunStatusApp(t *testing.T) {
	server.Start()
	r := api.NewAppManager().Status("2a7d6def46b64e059a869abdfe662990")
	fmt.Println(r.Data)
}

// 滚动日志读取
func TestReadLog(t *testing.T) {
	logPath := "logs/app/adpublish_0d59d93a0d9947b4a8d455c04f92e156.log"
	file, err := os.Open(logPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	// 设置一个偏移量，表示我们从文件的末尾开始读取
	_, err = file.Seek(-1024, io.SeekEnd)
	br := bufio.NewReader(file)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		fmt.Println(string(a))
	}
	file.Close()
}

func TestAppSelect(t *testing.T) {
	server.Start()
	ids := []string{"0d59d93a0d9947b4a8d455c04f92e156", "ee", "ff"}
	var apps []*model.App
	err := global.DB.Find(&apps, "id in (?)", ids).Error
	if err != nil {
		fmt.Println(err)
	}
	for _, app := range apps {
		fmt.Println(app)
	}
}

func TestSettingSave(t *testing.T) {
	b := false
	server.Start()
	r := api.NewGlobalSettingManager().SetGlobalSetting(model.GlobalSetting{
		DefaultWeb:     "http://www.baidu.com",
		OpenDefaultWeb: &b,
	})
	fmt.Println(r)
}

func TestProcKill(t *testing.T) {
	process, err := os.FindProcess(7380)
	if err != nil {
		return
	}
	fmt.Println(process.Pid)
	pid := uint32(7380) // 替换为你要杀死的进程组的主进程ID
	err = osutils.KillProcess([]uint32{pid})
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("进程组已成功终止")
	}
}

func TestFindPids(t *testing.T) {
	pids, err := osutils.ProcGroupPids(7380)
	if err != nil {
		return
	}
	fmt.Println(pids)
}

func TestRefreshPorts(t *testing.T) {
	server.Start()
	r := api.NewAppManager().RefreshPorts("fb9550a9a9344cf79a27035c07cd7bff")
	fmt.Println(r.Data)
}
