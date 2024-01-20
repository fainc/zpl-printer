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
