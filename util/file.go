package util

import (
	"os"
)

// FileIsExist 确认文件是否存在
func FileIsExist(file string) bool {
	if _, err := os.Stat(file); os.IsExist(err) || err == nil {
		return true
	}	
	return false
}