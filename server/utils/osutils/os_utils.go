package osutils

import (
	"fmt"
	"github.com/jaypipes/ghw"
	"github.com/jaypipes/ghw/pkg/baseboard"
	"github.com/jaypipes/ghw/pkg/bios"
	"github.com/jaypipes/ghw/pkg/block"
	"github.com/jaypipes/ghw/pkg/chassis"
	"github.com/jaypipes/ghw/pkg/cpu"
	"github.com/jaypipes/ghw/pkg/gpu"
	"github.com/jaypipes/ghw/pkg/memory"
	"github.com/jaypipes/ghw/pkg/net"
	"github.com/jaypipes/ghw/pkg/product"
	"github.com/jaypipes/ghw/pkg/topology"
	"github.com/konglingyinxia/win-start-client/server/dto"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// DiskParse Disk 磁盘统计信息
func DiskParse() dto.DiskCount {
	partitions, err := disk.Partitions(false)
	if err != nil {
		return dto.DiskCount{}
	}
	var total uint64
	var used uint64
	for _, partition := range partitions {
		usage, _ := disk.Usage(partition.Mountpoint)
		total += usage.Total
		used += usage.Used
	}
	return dto.DiskCount{
		Total:       total,
		Used:        used,
		UsedPercent: float64(used) / float64(total) * 100,
		Free:        total - used,
	}
}

// DirSize 磁盘路径统计信息
func DirSize(dir string) int64 {
	var totalSize int64
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		totalSize += info.Size()
		return nil
	})
	if err != nil {
		return 0
	}
	return totalSize
}

func MemParse() interface{} {
	//获取物理内存和交换区内存信息
	m1, _ := mem.VirtualMemory()
	m2, _ := mem.SwapMemory()
	if m1 != nil && m2 != nil {
		totalMemory := strconv.FormatFloat(float64(m1.Total)/1024/1024/1024.0, 'f', 2, 64)
		usedMemory := strconv.FormatFloat(float64(m1.Used)/1024/1024/1024.0, 'f', 2, 64)
		percent := strconv.FormatFloat(m1.UsedPercent, 'f', 2, 64)
		return map[string]string{
			"totalMemory": totalMemory,
			"usedMemory":  usedMemory,
			"percent":     percent,
		}
	}
	return map[string]string{}
}
func GpuParse(gpus *gpu.Info) string {
	result := ""
	for _, value := range gpus.GraphicsCards {
		v := value.DeviceInfo.Vendor
		if strings.HasPrefix(v.Name, "NVIDIA") {
			result = fmt.Sprintf("%s", value.DeviceInfo.Product.Name)
			break
		}
	}
	if result == "" {
		result = "GPU:Unknown"
	}
	return result
}

func Host(opts ...*ghw.WithOption) (*ghw.HostInfo, error) {
	memInfo, err := memory.New(opts...)
	if err != nil {
		return nil, err
	}
	blockInfo, err := block.New(opts...)
	if err != nil {
		return nil, err
	}
	cpuInfo, err := cpu.New(opts...)
	if err != nil {
		return nil, err
	}
	topologyInfo, err := topology.New(opts...)
	if err != nil {
		return nil, err
	}
	netInfo, err := net.New(opts...)
	if err != nil {
		return nil, err
	}
	gpuInfo, err := gpu.New(opts...)
	if err != nil {
		return nil, err
	}
	chassisInfo, err := chassis.New(opts...)
	if err != nil {
		return nil, err
	}
	biosInfo, err := bios.New(opts...)
	if err != nil {
		return nil, err
	}
	baseboardInfo, err := baseboard.New(opts...)
	if err != nil {
		return nil, err
	}
	productInfo, err := product.New(opts...)
	if err != nil {
		return nil, err
	}
	return &ghw.HostInfo{
		CPU:       cpuInfo,
		Memory:    memInfo,
		Block:     blockInfo,
		Topology:  topologyInfo,
		Network:   netInfo,
		GPU:       gpuInfo,
		Chassis:   chassisInfo,
		BIOS:      biosInfo,
		Baseboard: baseboardInfo,
		Product:   productInfo,
	}, nil
}

// StatusWithContext 获取进程状态
func StatusWithContext(pid int32) {

}
