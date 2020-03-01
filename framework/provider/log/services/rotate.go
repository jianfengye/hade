package services

import (
	"context"
	pkgLog "log"
	"path/filepath"

	"github.com/jianfengye/hade/framework/contract"
	"github.com/jianfengye/hade/framework/provider/log/formatter"
	"github.com/natefinch/lumberjack"
)

type HadeRotateLog struct {
	HadeLog

	folder     string
	file       string
	maxFiles   int
	dateFormat string

	rlogger *lumberjack.Logger
}

func NewHadeRotateLog(params ...interface{}) (interface{}, error) {
	level := params[0].(contract.LogLevel)
	ctxFielder := params[1].(contract.CtxFielder)
	formatter := params[2].(contract.Formatter)
	configs := params[3].(map[string]interface{})
	folder := configs["folder"].(string)
	file := configs["file"].(string)
	maxFiles := configs["max_files"].(int)
	dateFormat := configs["date_format"].(string)

	log := &HadeRotateLog{}
	log.SetLevel(level)
	log.SetCxtFielder(ctxFielder)
	log.SetFormatter(formatter)
	log.SetFile(file)
	log.SetFolder(folder)
	log.SetMaxFiles(maxFiles)
	log.SetDateFormat(dateFormat)

	log.rlogger = &lumberjack.Logger{
		Filename: filepath.Join(log.folder, log.file),
		MaxAge:   log.maxFiles,
		MaxSize:  1000,
	}
	return log, nil
}

func (l *HadeRotateLog) SetFolder(folder string) {
	l.folder = folder
}

func (l *HadeRotateLog) SetFile(file string) {
	l.file = file
}

func (l *HadeRotateLog) SetMaxFiles(maxFiles int) {
	l.maxFiles = maxFiles
}

func (l *HadeRotateLog) SetDateFormat(dateFormat string) {
	l.dateFormat = dateFormat
}

func (l *HadeRotateLog) IsLevelEnable(level contract.LogLevel) bool {
	return level <= l.level
}

func (log *HadeRotateLog) logf(level contract.LogLevel, ctx context.Context, msg string, fields []interface{}) error {
	if !log.IsLevelEnable(level) {
		return nil
	}
	prefix := ""
	switch level {
	case contract.PanicLevel:
		prefix = "[Panic]"
	case contract.FatalLevel:
		prefix = "[Fatal]"
	case contract.ErrorLevel:
		prefix = "[Error]"
	case contract.WarnLevel:
		prefix = "[Warn]"
	case contract.InfoLevel:
		prefix = "[Info]"
	case contract.DebugLevel:
		prefix = "[Debug]"
	case contract.TraceLevel:
		prefix = "[Trace]"
	}

	fs := fields
	if log.ctxFielder != nil {
		t := log.ctxFielder(ctx)
		if t != nil {
			fs = append(fs, t...)
		}
	}
	if log.formatter == nil {
		log.formatter = formatter.TextFormatter
	}
	ct, err := log.formatter(msg, fs)
	if err != nil {
		return err
	}

	if level == contract.PanicLevel {
		pkgLog.Panic(ct)
		return nil
	}

	log.rlogger.Write([]byte(prefix))
	log.rlogger.Write(ct)
	log.rlogger.Write([]byte{'\n'})
	return nil
}

// Panic will call panic(fields) for debug
func (log *HadeRotateLog) Panic(ctx context.Context, msg string, fields []interface{}) {
	log.logf(contract.PanicLevel, ctx, msg, fields)
}

// Fatal will add fatal record which contains msg and fields
func (log *HadeRotateLog) Fatal(ctx context.Context, msg string, fields []interface{}) {
	log.logf(contract.FatalLevel, ctx, msg, fields)
}

// Error will add error record which contains msg and fields
func (log *HadeRotateLog) Error(ctx context.Context, msg string, fields []interface{}) {
	log.logf(contract.ErrorLevel, ctx, msg, fields)
}

// Warn will add warn record which contains msg and fields
func (log *HadeRotateLog) Warn(ctx context.Context, msg string, fields []interface{}) {
	log.logf(contract.WarnLevel, ctx, msg, fields)
}

// Info will add info record which contains msg and fields
func (log *HadeRotateLog) Info(ctx context.Context, msg string, fields []interface{}) {
	log.logf(contract.InfoLevel, ctx, msg, fields)
}

// Debug will add debug record which contains msg and fields
func (log *HadeRotateLog) Debug(ctx context.Context, msg string, fields []interface{}) {
	log.logf(contract.DebugLevel, ctx, msg, fields)
}

// Trace will add trace info which contains msg and fields
func (log *HadeRotateLog) Trace(ctx context.Context, msg string, fields []interface{}) {
	log.logf(contract.TraceLevel, ctx, msg, fields)
}
