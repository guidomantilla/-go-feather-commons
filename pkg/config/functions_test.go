package config

import (
	"context"
	"testing"

	feather_commons_environment "github.com/guidomantilla/go-feather-commons/pkg/environment"
)

func TestProcess(t *testing.T) {
	type args struct {
		ctx         context.Context
		environment feather_commons_environment.Environment
		config      interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Process(tt.args.ctx, tt.args.environment, tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
