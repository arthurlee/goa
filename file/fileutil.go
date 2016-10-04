package fileutil

import (
	"os"
)

// check if path is exists
func Exist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// check if path is a directory
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// check if path is a file
func IsFile(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.Mode().IsRegular()
}
