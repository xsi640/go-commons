package os

import "runtime"

func Name() string {
	return runtime.GOOS
}

func Arch() string {
	return runtime.GOARCH
}

func Is64bit() bool {
	return uint64(^uintptr(0)) == ^uint64(0)
}
