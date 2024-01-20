package zpl_printer

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

func String2UintPTR(s string) (ret uintptr) {
	fromString, err := windows.UTF16PtrFromString(s)
	if err != nil {
		panic(err.Error())
	}
	ret = uintptr(unsafe.Pointer(fromString))
	return
}

func StatusDecode(status uintptr) int {
	if status == 0 {
		return 0
	}
	if (status & 0b1) > 0 {
		return 1
	}
	if (status & 0b10) > 0 {
		return 2
	}
	if (status & 0b100) > 0 {
		return 3
	}
	if (status & 0b1000) > 0 {
		return 4
	}
	if (status & 0b10000) > 0 {
		return 5
	}
	if (status & 0b100000) > 0 {
		return 6
	}
	if (status & 0b1000000) > 0 {
		return 7
	}
	return 8
}
