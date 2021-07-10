// +build darwin

package main

/*
#include <mach/mach_host.h>
*/
import "C"

import (
	"encoding/binary"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Helper function to get hardware memory size
func getHwMemsize() (uint64, error) {
	totalString, err := unix.Sysctl("hw.memsize")
	if err != nil {
		return 0, err
	}

	totalString += "\x00"
	total := uint64(binary.LittleEndian.Uint64([]byte(totalString)))

	return total, nil
}

// getMemoryUsage returns used memory in percentage
func getMemoryUsage() (float64, error) {
	count := C.mach_msg_type_number_t(C.HOST_VM_INFO_COUNT)
	var vmstat C.vm_statistics_data_t

	C.host_statistics(C.host_t(C.mach_host_self()),
		C.HOST_VM_INFO,
		C.host_info_t(unsafe.Pointer(&vmstat)),
		&count)

	pageSize := uint64(unix.Getpagesize())
	total, err := getHwMemsize()
	if err != nil {
		return 0, err
	}
	totalCount := C.natural_t(total / pageSize)

	availableCount := vmstat.inactive_count + vmstat.free_count
	usedPercent := 100 * float64(totalCount-availableCount) / float64(totalCount)

	return usedPercent, nil
}
