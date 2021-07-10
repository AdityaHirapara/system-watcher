// +build windows

package main

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

var Modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
var procGlobalMemoryStatusEx = Modkernel32.NewProc("GlobalMemoryStatusEx")

type memoryStatusEx struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

func getMemoryUsage() (float64, error) {
	var memInfo memoryStatusEx
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	mem, _, _ := procGlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	if mem == 0 {
		return 0, windows.GetLastError()
	}

	ret := float64(memInfo.dwMemoryLoad)

	return ret, nil
}
