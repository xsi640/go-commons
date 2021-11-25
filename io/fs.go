package fs

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Mkdir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func Delete(path string) error {
	return os.RemoveAll(path)
}

func Move(src, dst string) error {
	return os.Rename(src, dst)
}

func CopyFile(src, dst string, override bool) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	if !override && Exists(dst) {
		return
	}

	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		if e := out.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	err = out.Sync()
	if err != nil {
		return
	}

	si, err := os.Stat(src)
	if err != nil {
		return
	}
	err = os.Chmod(dst, si.Mode())
	if err != nil {
		return
	}
	return
}

func CopyDir(src string, dst string, override bool) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	err = os.MkdirAll(dst, si.Mode())
	if err != nil {
		return
	}

	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath, override)
			if err != nil {
				return
			}
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath, override)
			if err != nil {
				return
			}
		}
	}

	return
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

func IsDir(path string) (bool, error) {
	f, err := os.Stat(path)
	if err == nil {
		return f.IsDir(), nil
	}
	return false, err
}

func IsFile(path string) (bool, error) {
	r, err := IsDir(path)
	return !r, err
}
