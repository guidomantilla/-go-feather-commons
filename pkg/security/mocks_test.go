// Code generated by MockGen. DO NOT EDIT.
// Source: ../pkg/security/types.go

// Package security is a generated GoMock package.
package security

import (
	reflect "reflect"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestNewMockPasswordEncoder(t *testing.T) {
	type args struct {
		ctrl *gomock.Controller
	}
	tests := []struct {
		name string
		args args
		want *MockPasswordEncoder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMockPasswordEncoder(tt.args.ctrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockPasswordEncoder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordEncoder_EXPECT(t *testing.T) {
	tests := []struct {
		name string
		m    *MockPasswordEncoder
		want *MockPasswordEncoderMockRecorder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.EXPECT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordEncoder.EXPECT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordEncoder_Encode(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		m       *MockPasswordEncoder
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Encode(tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockPasswordEncoder.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordEncoder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordEncoderMockRecorder_Encode(t *testing.T) {
	type args struct {
		rawPassword interface{}
	}
	tests := []struct {
		name string
		mr   *MockPasswordEncoderMockRecorder
		args args
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.Encode(tt.args.rawPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordEncoderMockRecorder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordEncoder_Matches(t *testing.T) {
	type args struct {
		encodedPassword string
		rawPassword     string
	}
	tests := []struct {
		name    string
		m       *MockPasswordEncoder
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Matches(tt.args.encodedPassword, tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockPasswordEncoder.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordEncoder.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordEncoderMockRecorder_Matches(t *testing.T) {
	type args struct {
		encodedPassword interface{}
		rawPassword     interface{}
	}
	tests := []struct {
		name string
		mr   *MockPasswordEncoderMockRecorder
		args args
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.Matches(tt.args.encodedPassword, tt.args.rawPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordEncoderMockRecorder.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordEncoder_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		m       *MockPasswordEncoder
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.UpgradeEncoding(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockPasswordEncoder.UpgradeEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordEncoder.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordEncoderMockRecorder_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword interface{}
	}
	tests := []struct {
		name string
		mr   *MockPasswordEncoderMockRecorder
		args args
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.UpgradeEncoding(tt.args.encodedPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordEncoderMockRecorder.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMockPasswordGenerator(t *testing.T) {
	type args struct {
		ctrl *gomock.Controller
	}
	tests := []struct {
		name string
		args args
		want *MockPasswordGenerator
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMockPasswordGenerator(tt.args.ctrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockPasswordGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordGenerator_EXPECT(t *testing.T) {
	tests := []struct {
		name string
		m    *MockPasswordGenerator
		want *MockPasswordGeneratorMockRecorder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.EXPECT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordGenerator.EXPECT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordGenerator_Generate(t *testing.T) {
	tests := []struct {
		name string
		m    *MockPasswordGenerator
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Generate(); got != tt.want {
				t.Errorf("MockPasswordGenerator.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordGeneratorMockRecorder_Generate(t *testing.T) {
	tests := []struct {
		name string
		mr   *MockPasswordGeneratorMockRecorder
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.Generate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordGeneratorMockRecorder.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordGenerator_Validate(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		m       *MockPasswordGenerator
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Validate(tt.args.rawPassword); (err != nil) != tt.wantErr {
				t.Errorf("MockPasswordGenerator.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMockPasswordGeneratorMockRecorder_Validate(t *testing.T) {
	type args struct {
		rawPassword interface{}
	}
	tests := []struct {
		name string
		mr   *MockPasswordGeneratorMockRecorder
		args args
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.Validate(tt.args.rawPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordGeneratorMockRecorder.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMockPasswordManager(t *testing.T) {
	type args struct {
		ctrl *gomock.Controller
	}
	tests := []struct {
		name string
		args args
		want *MockPasswordManager
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMockPasswordManager(tt.args.ctrl); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMockPasswordManager() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManager_EXPECT(t *testing.T) {
	tests := []struct {
		name string
		m    *MockPasswordManager
		want *MockPasswordManagerMockRecorder
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.EXPECT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManager.EXPECT() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManager_Encode(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		m       *MockPasswordManager
		args    args
		want    *string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Encode(tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockPasswordManager.Encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManager.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManagerMockRecorder_Encode(t *testing.T) {
	type args struct {
		rawPassword interface{}
	}
	tests := []struct {
		name string
		mr   *MockPasswordManagerMockRecorder
		args args
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.Encode(tt.args.rawPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManagerMockRecorder.Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManager_Generate(t *testing.T) {
	tests := []struct {
		name string
		m    *MockPasswordManager
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.m.Generate(); got != tt.want {
				t.Errorf("MockPasswordManager.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManagerMockRecorder_Generate(t *testing.T) {
	tests := []struct {
		name string
		mr   *MockPasswordManagerMockRecorder
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.Generate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManagerMockRecorder.Generate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManager_Matches(t *testing.T) {
	type args struct {
		encodedPassword string
		rawPassword     string
	}
	tests := []struct {
		name    string
		m       *MockPasswordManager
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.Matches(tt.args.encodedPassword, tt.args.rawPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockPasswordManager.Matches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManager.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManagerMockRecorder_Matches(t *testing.T) {
	type args struct {
		encodedPassword interface{}
		rawPassword     interface{}
	}
	tests := []struct {
		name string
		mr   *MockPasswordManagerMockRecorder
		args args
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.Matches(tt.args.encodedPassword, tt.args.rawPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManagerMockRecorder.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManager_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword string
	}
	tests := []struct {
		name    string
		m       *MockPasswordManager
		args    args
		want    *bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.m.UpgradeEncoding(tt.args.encodedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockPasswordManager.UpgradeEncoding() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManager.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManagerMockRecorder_UpgradeEncoding(t *testing.T) {
	type args struct {
		encodedPassword interface{}
	}
	tests := []struct {
		name string
		mr   *MockPasswordManagerMockRecorder
		args args
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.UpgradeEncoding(tt.args.encodedPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManagerMockRecorder.UpgradeEncoding() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockPasswordManager_Validate(t *testing.T) {
	type args struct {
		rawPassword string
	}
	tests := []struct {
		name    string
		m       *MockPasswordManager
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.Validate(tt.args.rawPassword); (err != nil) != tt.wantErr {
				t.Errorf("MockPasswordManager.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMockPasswordManagerMockRecorder_Validate(t *testing.T) {
	type args struct {
		rawPassword interface{}
	}
	tests := []struct {
		name string
		mr   *MockPasswordManagerMockRecorder
		args args
		want *gomock.Call
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mr.Validate(tt.args.rawPassword); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MockPasswordManagerMockRecorder.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
