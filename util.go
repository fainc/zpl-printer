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

// StatusDecode 按位与操作后以新 int 值形式直接表示文档对应二进制值，方便理解和后续逻辑判断
func StatusDecode(status uintptr) int {
	if status == 0 { // 正常
		return 0
	}
	if (status & 0b1) > 0 { // 打印头被打开
		return 1
	}
	if (status & 0b10) > 0 { // 卡纸
		return 2
	}
	if (status & 0b100) > 0 { // 缺纸
		return 4
	}
	if (status & 0b1000) > 0 { // 缺碳带
		return 8
	}
	if (status & 0b10000) > 0 { // 打印暂停
		return 16
	}
	if (status & 0b100000) > 0 { // 打印中
		return 32
	}
	if (status & 0b1000000) > 0 { // 上盖打开
		return 64
	}
	return 128 // 其它错误
}
