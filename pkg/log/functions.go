package log

import (
	"context"
	"os"
	"sync/atomic"
)

var singleton atomic.Value

func Default() Logger {
	value := singleton.Load()
	if value == nil {
		logger := NewDefaultLogger()
		singleton.Store(logger)
		return logger
	}
	return value.(Logger)
}

func Debug(msg string, args ...any) {
	slogLogger := Default()
	slogLogger.Debug(context.Background(), msg, args...)
}

func Info(msg string, args ...any) {
	slogLogger := Default()
	slogLogger.Info(context.Background(), msg, args...)
}

func Warn(msg string, args ...any) {
	slogLogger := Default()
	slogLogger.Warn(context.Background(), msg, args...)
}

func Error(msg string, args ...any) {
	slogLogger := Default()
	slogLogger.Error(context.Background(), msg, args...)
}

func Fatal(msg string, args ...any) {
	slogLogger := Default()
	slogLogger.Fatal(context.Background(), msg, args...)
	os.Exit(1)
}
