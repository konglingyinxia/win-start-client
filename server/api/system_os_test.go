package api

import (
	"fmt"
	"github.com/jaypipes/ghw"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"testing"
)

func TestSystemOs(t *testing.T) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return
	}
	fmt.Println(partitions)

	usage, err := disk.Usage("/")
	if err != nil {
		return
	}
	fmt.Println(usage.String())

}
func TestHost(t *testing.T) {
	info, err := host.Info()
	fmt.Println(info)
	hostInfo, err := ghw.Host()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hostInfo)

}
