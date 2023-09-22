package config

import (
	"context"

	envconfig "github.com/sethvargo/go-envconfig"

	feather_commons_environment "github.com/guidomantilla/go-feather-commons/pkg/environment"
)

func Process(ctx context.Context, environment feather_commons_environment.Environment, config interface{}) error {
	return envconfig.ProcessWith(ctx, config, &EnvironmentLookup{
		environment: environment,
	})
}
