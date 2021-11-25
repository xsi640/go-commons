package paths

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

const PathSeparator = os.PathSeparator

func Join(elem ...string) string {
	return filepath.Join(elem...)
}

func ExtName(file string) string {
	return filepath.Ext(file)
}

func WithoutExtName(file string) string {
	filename := filepath.Base(file)
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return filename
	} else {
		return filename[0:index]
	}
}

func Parent(file string) string {
	return path.Dir(file)
}
