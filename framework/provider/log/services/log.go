package services

import (
	"context"
	pkgLog "log"

	"github.com/jianfengye/hade/framework/contract"
	"github.com/jianfengye/hade/framework/provider/log/formatter"
)

type HadeLog struct {
	level      contract.LogLevel
	formatter  contract.Formatter
	ctxFielder contract.CtxFielder

	logger *pkgLog.Logger
}

func (log *HadeLog) IsLevelEnable(level contract.LogLevel) bool {
	return level <= log.level
}

func (log *HadeLog) logf(logger *pkgLog.Logger, level contract.LogLevel, ctx context.Context, msg string, fields []interface{}) error {
	if !log.IsLevelEnable(level) {
		return nil
	}
	prefix := ""
	switch level {
	case contract.PanicLevel:
		prefix = "[Panic] "
	case contract.FatalLevel:
		prefix = "[Fatal] "
	case contract.ErrorLevel:
		prefix = "[Error] "
	case contract.WarnLevel:
		prefix = "[Warn] "
	case contract.InfoLevel:
		prefix = "[Info] "
	case contract.DebugLevel:
		prefix = "[Debug] "
	case contract.TraceLevel:
		prefix = "[Trace] "
	}
	logger.SetPrefix(prefix)
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
		logger.Panicln(string(ct))
		return nil
	}

	logger.Println(string(ct))
	return nil
}

// Panic will call panic(fields) for debug
func (log *HadeLog) Panic(ctx context.Context, msg string, fields []interface{}) {
	log.logf(log.logger, contract.PanicLevel, ctx, msg, fields)
}

// Fatal will add fatal record which contains msg and fields
func (log *HadeLog) Fatal(ctx context.Context, msg string, fields []interface{}) {
	log.logf(log.logger, contract.FatalLevel, ctx, msg, fields)
}

// Error will add error record which contains msg and fields
func (log *HadeLog) Error(ctx context.Context, msg string, fields []interface{}) {
	log.logf(log.logger, contract.ErrorLevel, ctx, msg, fields)
}

// Warn will add warn record which contains msg and fields
func (log *HadeLog) Warn(ctx context.Context, msg string, fields []interface{}) {
	log.logf(log.logger, contract.WarnLevel, ctx, msg, fields)
}

// Info will add info record which contains msg and fields
func (log *HadeLog) Info(ctx context.Context, msg string, fields []interface{}) {
	log.logf(log.logger, contract.InfoLevel, ctx, msg, fields)
}

// Debug will add debug record which contains msg and fields
func (log *HadeLog) Debug(ctx context.Context, msg string, fields []interface{}) {
	log.logf(log.logger, contract.DebugLevel, ctx, msg, fields)
}

// Trace will add trace info which contains msg and fields
func (log *HadeLog) Trace(ctx context.Context, msg string, fields []interface{}) {
	log.logf(log.logger, contract.TraceLevel, ctx, msg, fields)
}

// SetLevel set log level, and higher level will be recorded
func (log *HadeLog) SetLevel(level contract.LogLevel) {
	log.level = level
}

// SetCxtFielder will get fields from context
func (log *HadeLog) SetCxtFielder(handler contract.CtxFielder) {
	log.ctxFielder = handler
}

// SetFormatter will set formatter handler will covert data to string for recording
func (log *HadeLog) SetFormatter(formatter contract.Formatter) {
	log.formatter = formatter
}
