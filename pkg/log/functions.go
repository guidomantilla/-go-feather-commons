package log

import (
	"context"
	"os"
	"strings"
	"sync/atomic"
)

var singleton atomic.Value

func retrieveSingleton() Logger {
	value := singleton.Load()
	if value == nil {
		return Default()
	}
	return value.(Logger)
}

func Default() Logger {
	level := UndefinedLoggerLevel.ValueFromName(strings.ToUpper("info"))
	format := UndefinedLoggerFormat.ValueFromName(strings.ToUpper("text"))
	logger := NewDefaultLogger(level, format)
	singleton.Store(logger)
	return logger
}

func Custom() Logger {
	logLevel, logFormat := os.Getenv("LOG_LEVEL"), os.Getenv("LOG_FORMAT")
	level := UndefinedLoggerLevel.ValueFromName(strings.ToUpper(logLevel))
	format := UndefinedLoggerFormat.ValueFromName(strings.ToUpper(logFormat))
	logger := NewDefaultLogger(level, format)
	singleton.Store(logger)
	return logger
}

func Debug(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Debug(context.Background(), msg, args...)
}

func Info(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Info(context.Background(), msg, args...)
}

func Warn(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Warn(context.Background(), msg, args...)
}

func Error(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Error(context.Background(), msg, args...)
}

func Fatal(msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Fatal(context.Background(), msg, args...)
	os.Exit(1)
}
