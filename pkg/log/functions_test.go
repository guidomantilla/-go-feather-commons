package log

import (
	"reflect"
	"testing"
)

func Test_retrieveSingleton(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := retrieveSingleton(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("retrieveSingleton() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefault(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Default(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Default() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustom(t *testing.T) {
	tests := []struct {
		name string
		want Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Custom(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Custom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDebug(t *testing.T) {
	type args struct {
		msg  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Debug(tt.args.msg, tt.args.args...)
		})
	}
}

func TestInfo(t *testing.T) {
	type args struct {
		msg  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.args.msg, tt.args.args...)
		})
	}
}

func TestWarn(t *testing.T) {
	type args struct {
		msg  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.args.msg, tt.args.args...)
		})
	}
}

func TestError(t *testing.T) {
	type args struct {
		msg  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.args.msg, tt.args.args...)
		})
	}
}

func TestFatal(t *testing.T) {
	type args struct {
		msg  string
		args []any
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Fatal(tt.args.msg, tt.args.args...)
		})
	}
}
