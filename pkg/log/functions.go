package log

import (
	"context"
	"io"
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

func Default(writers ...io.Writer) Logger {
	logger := NewSlogLogger(SlogLevelOff.ValueFromName("INFO"), writers...)
	singleton.Store(logger)
	return logger
}

func Custom(writers ...io.Writer) Logger {
	logger := NewSlogLogger(SlogLevelOff.ValueFromName(strings.ToUpper(os.Getenv("LOG_LEVEL"))), writers...)
	singleton.Store(logger)
	return logger
}

//

func Debug(ctx context.Context, msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Debug(ctx, msg, args...)
}

func Info(ctx context.Context, msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Info(ctx, msg, args...)
}

func Warn(ctx context.Context, msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Warn(ctx, msg, args...)
}

func Error(ctx context.Context, msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Error(ctx, msg, args...)
}

func Fatal(ctx context.Context, msg string, args ...any) {
	slogLogger := retrieveSingleton()
	slogLogger.Fatal(ctx, msg, args...)
	os.Exit(1)
}
