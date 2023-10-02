package resilience

import (
	"reflect"
	"testing"
)

func TestNewNoFailSaver(t *testing.T) {
	tests := []struct {
		name string
		want *NoFailSaver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewNoFailSaver(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNoFailSaver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoFailSaver_Run(t *testing.T) {
	type args struct {
		target func() error
	}
	tests := []struct {
		name     string
		failSafe *NoFailSaver
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.failSafe.Run(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("NoFailSaver.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
