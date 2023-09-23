package log

import (
	"context"
	"log/slog"
	"os"
)

const (
	SlogLevelDebug   = slog.LevelDebug
	SlogLevelInfo    = slog.LevelInfo
	SlogLevelWarning = slog.LevelWarn
	SlogLevelError   = slog.LevelError
	SlogLevelFatal   = slog.Level(12)
	SlogLevelOff     = slog.Level(16)
)

type SlogLogger struct {
	internal *slog.Logger
}

func NewDefaultLogger() *SlogLogger {
	opts := &slog.HandlerOptions{
		Level: SlogLevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)

				// This could also look up the name from a map or other structure, but
				// this demonstrates using a switch statement to rename levels. For
				// maximum performance, the string values should be constants, but this
				// example uses the raw strings for readability.
				switch {
				case level == SlogLevelDebug:
					a.Value = slog.StringValue("DEBUG")
				case level == SlogLevelInfo:
					a.Value = slog.StringValue("INFO")
				case level == SlogLevelWarning:
					a.Value = slog.StringValue("WARN")
				case level == SlogLevelError:
					a.Value = slog.StringValue("ERROR")
				case level == SlogLevelFatal:
					a.Value = slog.StringValue("FATAL")
				case level == SlogLevelOff:
					a.Value = slog.StringValue("OFF")
				default:
					a.Value = slog.StringValue("INFO")
				}
			}

			return a
		},
	}
	internal := slog.New(slog.NewJSONHandler(os.Stdout, opts))
	slog.SetDefault(internal)
	slogLogger := &SlogLogger{internal: internal}
	singleton.Store(slogLogger)
	return slogLogger
}

func (logger *SlogLogger) Debug(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelDebug, msg, args...)
}

func (logger *SlogLogger) Info(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelInfo, msg, args...)
}

func (logger *SlogLogger) Warn(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelWarning, msg, args...)
}

func (logger *SlogLogger) Error(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelError, msg, args...)
}

func (logger *SlogLogger) Fatal(ctx context.Context, msg string, args ...any) {
	logger.internal.Log(ctx, SlogLevelFatal, msg, args...)
}
