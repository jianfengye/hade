package services

import (
	"io"
	pkgLog "log"
	"os"

	"github.com/jianfengye/hade/framework/contract"
)

type HadeConsoleLog struct {
	HadeLog
}

func NewHadeConsoleLog(params ...interface{}) (interface{}, error) {
	level := params[0].(contract.LogLevel)
	ctxFielder := params[1].(contract.CtxFielder)
	formatter := params[2].(contract.Formatter)

	log := &HadeConsoleLog{}

	log.SetLevel(level)
	log.SetCxtFielder(ctxFielder)
	log.SetFormatter(formatter)

	log.logger = pkgLog.New(os.Stdout, "", pkgLog.Ldate|pkgLog.Ltime)
	return log, nil
}

func (l *HadeConsoleLog) SetOutput(out io.Writer) {
	l.logger.SetOutput(out)
}
