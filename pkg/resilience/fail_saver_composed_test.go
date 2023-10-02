package resilience

import (
	"reflect"
	"testing"
)

func TestWithCircuitBreaker(t *testing.T) {
	type args struct {
		breaker FailSaver
	}
	tests := []struct {
		name string
		args args
		want DefaultFailSaverOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithCircuitBreaker(tt.args.breaker); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithCircuitBreaker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithRetrier(t *testing.T) {
	type args struct {
		retrier FailSaver
	}
	tests := []struct {
		name string
		args args
		want DefaultFailSaverOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithRetrier(tt.args.retrier); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithRetrier() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewDefaultFailSaver(t *testing.T) {
	type args struct {
		opts []DefaultFailSaverOption
	}
	tests := []struct {
		name string
		args args
		want *DefaultFailSaver
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDefaultFailSaver(tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDefaultFailSaver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultFailSaver_Run(t *testing.T) {
	type args struct {
		target func() error
	}
	tests := []struct {
		name     string
		failSafe *DefaultFailSaver
		args     args
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.failSafe.Run(tt.args.target); (err != nil) != tt.wantErr {
				t.Errorf("DefaultFailSaver.Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
