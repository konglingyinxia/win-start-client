package ip

import (
	"github.com/konglingyinxia/win-start-client/server/global"
	"net"
	"strings"
)

var localIp string

func LocalOutboundIp() (ip string) {
	if localIp != "" {
		return localIp
	}
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		global.LOG.Error("获取本机IP地址失败:" + err.Error())
		return "127.0.0.1"
	}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(addr.String(), ":")[0]
	localIp = ip
	return
}
