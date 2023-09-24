package log

import (
	"context"
	"io"
	"log/slog"
	"os"
)

type SlogLogger struct {
	internal *slog.Logger
}

func NewSlogLogger(level LoggerLevel, format LoggerFormat) *SlogLogger {
	opts := &slog.HandlerOptions{
		Level: SlogLevelOff.ValueFromLoggerLevel(level).ToSlogLevel(),
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				a.Value = slog.StringValue(SlogLevelOff.ValueFromSlogLevel(level).String())
			}
			return a
		},
	}
	internal := slog.New(format.SlogHandlerFunc(os.Stdout, opts))
	slog.SetDefault(internal)
	slogLogger := &SlogLogger{internal: internal}
	return slogLogger
}

func (logger *SlogLogger) Debug(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelDebug.ToSlogLevel(), msg, args...)
}

func (logger *SlogLogger) Info(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelInfo.ToSlogLevel(), msg, args...)
}

func (logger *SlogLogger) Warn(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelWarning.ToSlogLevel(), msg, args...)
}

func (logger *SlogLogger) Error(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelError.ToSlogLevel(), msg, args...)
}

func (logger *SlogLogger) Fatal(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelFatal.ToSlogLevel(), msg, args...)
	os.Exit(1)
}

func (logger *SlogLogger) RetrieveLogger() any {
	return logger.internal
}

//

const (
	SlogLevelDebug SlogLevel = iota
	SlogLevelInfo
	SlogLevelWarning
	SlogLevelError
	SlogLevelFatal
	SlogLevelOff
)

type SlogLevel int

func (enum SlogLevel) String() string {

	switch enum {
	case SlogLevelDebug:
		return "DEBUG"
	case SlogLevelInfo:
		return "INFO"
	case SlogLevelWarning:
		return "WARN"
	case SlogLevelError:
		return "ERROR"
	case SlogLevelFatal:
		return "FATAL"
	case SlogLevelOff:
		return "OFF"
	}
	return "OFF"
}

func (enum SlogLevel) ValueFromName(slogLevel string) SlogLevel {
	switch slogLevel {
	case "DEBUG":
		return SlogLevelDebug
	case "INFO":
		return SlogLevelInfo
	case "WARN":
		return SlogLevelWarning
	case "ERROR":
		return SlogLevelError
	case "FATAL":
		return SlogLevelFatal
	case "OFF":
		return SlogLevelOff
	}
	return SlogLevelOff
}

func (enum SlogLevel) ValueFromCardinal(slogLevel int) SlogLevel {
	switch slogLevel {
	case int(SlogLevelDebug):
		return SlogLevelDebug
	case int(SlogLevelInfo):
		return SlogLevelInfo
	case int(SlogLevelWarning):
		return SlogLevelWarning
	case int(SlogLevelError):
		return SlogLevelError
	case int(SlogLevelFatal):
		return SlogLevelFatal
	case int(SlogLevelOff):
		return SlogLevelOff
	}
	return SlogLevelOff
}

func (enum SlogLevel) ValueFromLoggerLevel(loggerLevel LoggerLevel) SlogLevel {
	switch loggerLevel {
	case DebugLoggerLevel:
		return SlogLevelDebug
	case InfoLoggerLevel:
		return SlogLevelInfo
	case WarnLoggerLevel:
		return SlogLevelWarning
	case ErrorLoggerLevel:
		return SlogLevelError
	case FatalLoggerLevel:
		return SlogLevelFatal
	case OffLoggerLevel:
		return SlogLevelOff
	}
	return SlogLevelOff
}

func (enum SlogLevel) ValueFromSlogLevel(slogLevel slog.Level) SlogLevel {
	switch slogLevel {
	case slog.LevelDebug:
		return SlogLevelDebug
	case slog.LevelInfo:
		return SlogLevelInfo
	case slog.LevelWarn:
		return SlogLevelWarning
	case slog.LevelError:
		return SlogLevelError
	case slog.Level(12):
		return SlogLevelFatal
	case slog.Level(16):
		return SlogLevelOff
	}
	return SlogLevelOff
}

func (enum SlogLevel) ToSlogLevel() slog.Level {
	switch enum {
	case SlogLevelDebug:
		return slog.LevelDebug
	case SlogLevelInfo:
		return slog.LevelInfo
	case SlogLevelWarning:
		return slog.LevelWarn
	case SlogLevelError:
		return slog.LevelError
	case SlogLevelFatal:
		return slog.Level(12)
	case SlogLevelOff:
		return slog.Level(16)
	}
	return slog.Level(16)
}

func (enum SlogLevel) ToLoggerLevel() LoggerLevel {
	switch enum {
	case SlogLevelDebug:
		return DebugLoggerLevel
	case SlogLevelInfo:
		return InfoLoggerLevel
	case SlogLevelWarning:
		return WarnLoggerLevel
	case SlogLevelError:
		return ErrorLoggerLevel
	case SlogLevelFatal:
		return FatalLoggerLevel
	case SlogLevelOff:
		return OffLoggerLevel
	}
	return OffLoggerLevel
}

//

func (enum LoggerFormat) SlogHandlerFunc(w io.Writer, opts *slog.HandlerOptions) slog.Handler {
	switch enum {
	case UndefinedLoggerFormat:
		return slog.NewTextHandler(w, opts)
	case TextLoggerFormat:
		return slog.NewTextHandler(w, opts)
	case JsonLoggerFormat:
		return slog.NewJSONHandler(w, opts)
	}
	return slog.NewTextHandler(w, opts)
}
