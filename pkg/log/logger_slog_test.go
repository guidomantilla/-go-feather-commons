package log

import (
	"bytes"
	"context"
	"log/slog"
	"reflect"
	"testing"
)

func TestNewSlogLogger(t *testing.T) {
	type args struct {
		level  LoggerLevel
		format LoggerFormat
	}
	tests := []struct {
		name string
		args args
		want *SlogLogger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSlogLogger(tt.args.level, tt.args.format); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLogger_Debug(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Debug(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_Info(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Info(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_Warn(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Warn(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_Error(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Error(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_Fatal(t *testing.T) {
	type args struct {
		ctx  context.Context
		msg  string
		args []any
	}
	tests := []struct {
		name   string
		logger *SlogLogger
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.logger.Fatal(tt.args.ctx, tt.args.msg, tt.args.args...)
		})
	}
}

func TestSlogLogger_RetrieveLogger(t *testing.T) {
	tests := []struct {
		name   string
		logger *SlogLogger
		want   any
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.logger.RetrieveLogger(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlogLogger.RetrieveLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLevel_String(t *testing.T) {
	tests := []struct {
		name string
		enum SlogLevel
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.String(); got != tt.want {
				t.Errorf("SlogLevel.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLevel_ValueFromName(t *testing.T) {
	type args struct {
		slogLevel string
	}
	tests := []struct {
		name string
		enum SlogLevel
		args args
		want SlogLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromName(tt.args.slogLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlogLevel.ValueFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLevel_ValueFromCardinal(t *testing.T) {
	type args struct {
		slogLevel int
	}
	tests := []struct {
		name string
		enum SlogLevel
		args args
		want SlogLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromCardinal(tt.args.slogLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlogLevel.ValueFromCardinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLevel_ValueFromLoggerLevel(t *testing.T) {
	type args struct {
		loggerLevel LoggerLevel
	}
	tests := []struct {
		name string
		enum SlogLevel
		args args
		want SlogLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromLoggerLevel(tt.args.loggerLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlogLevel.ValueFromLoggerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLevel_ValueFromSlogLevel(t *testing.T) {
	type args struct {
		slogLevel slog.Level
	}
	tests := []struct {
		name string
		enum SlogLevel
		args args
		want SlogLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromSlogLevel(tt.args.slogLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlogLevel.ValueFromSlogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLevel_ToSlogLevel(t *testing.T) {
	tests := []struct {
		name string
		enum SlogLevel
		want slog.Level
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ToSlogLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlogLevel.ToSlogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlogLevel_ToLoggerLevel(t *testing.T) {
	tests := []struct {
		name string
		enum SlogLevel
		want LoggerLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ToLoggerLevel(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SlogLevel.ToLoggerLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerFormat_SlogHandlerFunc(t *testing.T) {
	type args struct {
		opts *slog.HandlerOptions
	}
	tests := []struct {
		name  string
		enum  LoggerFormat
		args  args
		want  slog.Handler
		wantW string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if got := tt.enum.SlogHandlerFunc(w, tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerFormat.SlogHandlerFunc() = %v, want %v", got, tt.want)
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("LoggerFormat.SlogHandlerFunc() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
