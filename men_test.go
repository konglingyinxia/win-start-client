package main

import (
	"fmt"
	"runtime"
	"testing"
)

func Test_men(t *testing.T) {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Printf("Allocated memory: %v bytes\n", mem.Alloc)
	fmt.Printf("Total memory allocated and not yet freed: %v bytes\n", mem.TotalAlloc)
	fmt.Printf("Number of heap objects: %v\n", mem.HeapObjects)
}
