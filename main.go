package main

import (
	"github.com/jianfengye/hade/app/console"
	"github.com/jianfengye/hade/framework"
)

func main() {
	container := framework.NewHadeContainer()
	console.RunCommand(container)
}
