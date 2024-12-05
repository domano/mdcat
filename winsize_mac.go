package main

import (
	"golang.org/x/sys/unix"
	"unsafe"
)

type winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func getWinsize() (*winsize, error) {
	ws := &winsize{}
	retCode, _, errno := unix.Syscall(unix.SYS_IOCTL,
		uintptr(unix.Stdin),
		uintptr(unix.TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)))

	if int(retCode) == -1 {
		return nil, errno
	}
	return ws, nil
}
