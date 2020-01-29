package util

import(
	"testing"
)

func Test_Folder(t *testing.T) {
	folder := getCurrentFolder()

	t.Log(folder)

	t.Log(RootFolder())
	t.Log(StorageFolder())
	t.Log(LogFolder())
	t.Log(TesterFolder())

}