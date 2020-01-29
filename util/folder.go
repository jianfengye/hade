package util

import (
	"path"
	"runtime"
)

func getCurrentFolder() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

// RootFolder 获取根目录
func RootFolder() string {
	return path.Dir(getCurrentFolder())
}

// StorageFolder 获取Storage文件夹
func StorageFolder() string {
	return path.Join(RootFolder(), "storage")
}

// LogFolder 获取Storage/log文件夹
func LogFolder() string {
	return path.Join(StorageFolder(), "log")
}

// TesterFolder 获取Tester文件夹
func TesterFolder() string {
	return path.Join(RootFolder(), "tester")
}