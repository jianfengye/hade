package util

import "testing"

func Test_FileIsExist(t *testing.T) {
	file := "/Users/yejianfeng/Documents/workspace/huoding/backend/env.default.yaml"
	if FileIsExist(file) != true {
		t.Error("check file is exist error")
	}
}