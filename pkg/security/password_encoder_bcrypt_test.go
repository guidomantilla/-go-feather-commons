package security

import (
	"reflect"
	"testing"
)

func TestNewBCryptPasswordEncoder(t *testing.T) {
	type args struct {
		options []BCryptPasswordEncoderOption
	}
	tests := []struct {
		name string
		args args
		want *BCryptPasswordEncoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBCryptPasswordEncoder(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBCryptPasswordEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWithBCryptCost(t *testing.T) {
	type args struct {
		cost int
	}
	tests := []struct {
		name string
		args args
		want BCryptPasswordEncoderOption
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WithBCryptCost(tt.args.cost); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithBCryptCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBCryptPasswordEncoder_Encode(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		encoder *BCryptPasswordEncoder
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.encoder.Encode(tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("BCryptPasswordEncoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BCryptPasswordEncoder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBCryptPasswordEncoder_Matches(t *testing.T) {
	type args struct {
		encodedPassword string
		rawPassword     string
	}
	tests := []struct {
		name    string
		encoder *BCryptPasswordEncoder
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.encoder.Matches(tt.args.encodedPassword, tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("BCryptPasswordEncoder.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BCryptPasswordEncoder.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBCryptPasswordEncoder_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		encoder *BCryptPasswordEncoder
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.encoder.UpgradeEncoding(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("BCryptPasswordEncoder.UpgradeEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BCryptPasswordEncoder.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}
