package config

import (
	envconfig "github.com/sethvargo/go-envconfig"

	feather_commons_environment "github.com/guidomantilla/go-feather-commons/pkg/environment"
)

var (
	_ envconfig.Lookuper = (*EnvironmentLookup)(nil)
)

type EnvironmentLookup struct {
	environment feather_commons_environment.Environment
}

func (lookuper *EnvironmentLookup) Lookup(key string) (string, bool) {
	value := lookuper.environment.GetValue(key).AsString()
	return value, value != ""
}
