package util

import "encoding/json"

// PanicJson 快速打印出一个变量，并直接退出，加速调试
func PanicJson(a interface{}) {
	bs, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		panic(err)
	}
	panic(string(bs))
}
