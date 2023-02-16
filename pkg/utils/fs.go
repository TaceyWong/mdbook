package utils

import "os"

func CreateDirIfNotExist(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return err
}

// Remove file any way
func RemoveFile(path string) error {
	return os.Remove(path)
}

// Remove dir any way
func RemoveDir(path string) error {
	return os.RemoveAll(path)
}
