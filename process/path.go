package process

import (
	"os"
	"path/filepath"
)

func GetRunPath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
