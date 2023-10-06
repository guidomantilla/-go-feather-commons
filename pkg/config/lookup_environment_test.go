package config

import (
	"os"
	"testing"

	feather_commons_environment "github.com/guidomantilla/go-feather-commons/pkg/environment"
)

func TestEnvironmentLookup_Lookup(t *testing.T) {

	type args struct {
		key string
	}
	tests := []struct {
		name     string
		lookuper *EnvironmentLookup
		args     args
		want     string
		want1    bool
	}{
		{
			name: "EnvVar exists",
			lookuper: &EnvironmentLookup{
				environment: func() feather_commons_environment.Environment {
					if err := os.Setenv("SOME_ENV_VAR", "some-value"); err != nil {
						t.Errorf(err.Error())
					}
					return feather_commons_environment.Default()
				}(),
			},
			args: args{
				key: "SOME_ENV_VAR",
			},
			want:  "some-value",
			want1: true,
		},
		{
			name: "EnvVar not exists",
			lookuper: &EnvironmentLookup{
				environment: func() feather_commons_environment.Environment {
					return feather_commons_environment.Default()
				}(),
			},
			args: args{
				key: "SOME_ENV_VAR_",
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.lookuper.Lookup(tt.args.key)
			if got != tt.want {
				t.Errorf("EnvironmentLookup.Lookup() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("EnvironmentLookup.Lookup() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
