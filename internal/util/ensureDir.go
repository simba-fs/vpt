package util

import (
	"os"
)

// EnsureDir ensures the dir exist
func EnsureDir(dirPath string) (string, error) {
	stat, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		// mkdir
		os.Mkdir(dirPath, 0775)
	} else if !stat.IsDir() {
		// remove
		if err := os.Remove(dirPath); err != nil {
			return "", err
		}
		// mkdir
		os.Mkdir(dirPath, 0775)
	}

	return dirPath, nil
}
