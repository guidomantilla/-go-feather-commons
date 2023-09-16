package config

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"

	feather_commons_environment "github.com/guidomantilla/go-feather-commons/pkg/environment"
)

type EnvironmentLookup struct {
	environment feather_commons_environment.Environment
}

func (lookuper *EnvironmentLookup) Lookup(key string) (string, bool) {
	value := lookuper.environment.GetValue(key).AsString()
	return value, value != ""
}

func Process(ctx context.Context, environment feather_commons_environment.Environment, config interface{}) error {
	return envconfig.ProcessWith(ctx, config, &EnvironmentLookup{
		environment: environment,
	})
}
