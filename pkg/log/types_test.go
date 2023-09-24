package log

import (
	"reflect"
	"testing"
)

func TestLoggerFormat_String(t *testing.T) {
	tests := []struct {
		name string
		enum LoggerFormat
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.String(); got != tt.want {
				t.Errorf("LoggerFormat.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerFormat_ValueFromName(t *testing.T) {
	type args struct {
		loggerFormat string
	}
	tests := []struct {
		name string
		enum LoggerFormat
		args args
		want LoggerFormat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromName(tt.args.loggerFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerFormat.ValueFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerFormat_ValueFromCardinal(t *testing.T) {
	type args struct {
		loggerFormat int
	}
	tests := []struct {
		name string
		enum LoggerFormat
		args args
		want LoggerFormat
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromCardinal(tt.args.loggerFormat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerFormat.ValueFromCardinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerLevel_String(t *testing.T) {
	tests := []struct {
		name string
		enum LoggerLevel
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.String(); got != tt.want {
				t.Errorf("LoggerLevel.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerLevel_ValueFromName(t *testing.T) {
	type args struct {
		loggerLevel string
	}
	tests := []struct {
		name string
		enum LoggerLevel
		args args
		want LoggerLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromName(tt.args.loggerLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerLevel.ValueFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoggerLevel_ValueFromCardinal(t *testing.T) {
	type args struct {
		loggerLevel int
	}
	tests := []struct {
		name string
		enum LoggerLevel
		args args
		want LoggerLevel
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.enum.ValueFromCardinal(tt.args.loggerLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerLevel.ValueFromCardinal() = %v, want %v", got, tt.want)
			}
		})
	}
}
