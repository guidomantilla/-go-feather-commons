package log

import (
	"context"
	"strings"
)

const (
	UndefinedLoggerLevel LoggerLevel = iota
	DebugLoggerLevel
	InfoLoggerLevel
	WarnLoggerLevel
	ErrorLoggerLevel
	FatalLoggerLevel
	OffLoggerLevel
)

type LoggerLevel int

func (enum LoggerLevel) String() string {

	switch enum {
	case UndefinedLoggerLevel:
		return "undefined"
	case DebugLoggerLevel:
		return "debug"
	case InfoLoggerLevel:
		return "info"
	case WarnLoggerLevel:
		return "warn"
	case ErrorLoggerLevel:
		return "error"
	case FatalLoggerLevel:
		return "fatal"
	case OffLoggerLevel:
		return "off"
	}
	return "undefined"
}

func (enum LoggerLevel) ValueFromName(loggerLevel string) LoggerLevel {
	switch strings.ToLower(loggerLevel) {
	case "undefined":
		return UndefinedLoggerLevel
	case "debug":
		return DebugLoggerLevel
	case "info":
		return InfoLoggerLevel
	case "warn":
		return WarnLoggerLevel
	case "error":
		return ErrorLoggerLevel
	case "fatal":
		return FatalLoggerLevel
	case "off":
		return OffLoggerLevel
	}
	return UndefinedLoggerLevel
}

func (enum LoggerLevel) ValueFromCardinal(loggerLevel int) LoggerLevel {
	switch loggerLevel {
	case int(UndefinedLoggerLevel):
		return UndefinedLoggerLevel
	case int(DebugLoggerLevel):
		return DebugLoggerLevel
	case int(InfoLoggerLevel):
		return InfoLoggerLevel
	case int(WarnLoggerLevel):
		return WarnLoggerLevel
	case int(ErrorLoggerLevel):
		return ErrorLoggerLevel
	case int(FatalLoggerLevel):
		return FatalLoggerLevel
	case int(OffLoggerLevel):
		return OffLoggerLevel
	}
	return UndefinedLoggerLevel
}

type Logger interface {
	Debug(ctx context.Context, msg string, args ...any)
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Fatal(ctx context.Context, msg string, args ...any)
}
