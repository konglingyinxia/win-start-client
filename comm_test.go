package main

import (
	"fmt"
	"github.com/konglingyinxia/win-start-client/server"
	"github.com/konglingyinxia/win-start-client/server/api"
	"testing"
)

func TestOpenFIle(t *testing.T) {
	server.Start()
	res := api.NewCommManager().OpenFileExplorer("D:\\win-app-start\\opt\\app")
	fmt.Println(res)

}
