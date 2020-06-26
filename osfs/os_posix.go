// +build !plan9,!windows

package osfs

import (
	"os"

	"syscall"
)

func (f *file) Lock() error {
	f.m.Lock()
	defer f.m.Unlock()

	return syscall.Flock(int(f.File.Fd()), syscall.LOCK_EX)
}

func (f *file) Unlock() error {
	f.m.Lock()
	defer f.m.Unlock()

	return syscall.Flock(int(f.File.Fd()), syscall.LOCK_UN)
}

func rename(from, to string) error {
	return os.Rename(from, to)
}
