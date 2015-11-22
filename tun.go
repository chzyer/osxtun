package main

import (
	"errors"
	"os"
)

// #include "utun.c"
import "C"

func OpenUTun() (*os.File, error) {
	fd := C.utun_open()
	if fd < 0 {
		return nil, errors.New(C.GoString(C.strerror(-fd)))
	}
	return os.NewFile(uintptr(fd), "/dev/tun0"), nil
}
