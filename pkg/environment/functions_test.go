package environment

import (
	"reflect"
	"testing"

	"github.com/guidomantilla/go-feather-commons/pkg/properties"
)

func Test_retrieveSingleton(t *testing.T) {
	tests := []struct {
		name string
		want Environment
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
		want Environment
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
	type args struct {
		cmdArgsArray *[]string
	}
	tests := []struct {
		name string
		args args
		want Environment
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Custom(tt.args.cmdArgsArray); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Custom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	type args struct {
		property string
	}
	tests := []struct {
		name string
		args args
		want EnvVar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValue(tt.args.property); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetValueOrDefault(t *testing.T) {
	type args struct {
		property     string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want EnvVar
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetValueOrDefault(tt.args.property, tt.args.defaultValue); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValueOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPropertySources(t *testing.T) {
	tests := []struct {
		name string
		want []properties.PropertySource
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPropertySources(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPropertySources() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppendPropertySources(t *testing.T) {
	type args struct {
		propertySources []properties.PropertySource
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AppendPropertySources(tt.args.propertySources...)
		})
	}
}
