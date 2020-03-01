package services

import (
	pkgLog "log"
	"os"
)

type HadeConsoleLog struct {
	HadeLog
}

func NewHadeConsoleLog(params ...interface{}) (interface{}, error) {
	log := &HadeConsoleLog{}
	log.logger = pkgLog.New(os.Stdout, "", pkgLog.Ldate|pkgLog.Ltime)
	return log, nil
}
