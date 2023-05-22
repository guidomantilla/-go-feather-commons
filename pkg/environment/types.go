package environment

import (
	"strconv"

	feather_commons_properties "github.com/guidomantilla/go-feather-commons/pkg/properties"
)

var _ Environment = (*DefaultEnvironment)(nil)

type Environment interface {
	GetValue(property string) EnvVar
	GetValueOrDefault(property string, defaultValue string) EnvVar
	GetPropertySources() []feather_commons_properties.PropertySource
	AppendPropertySources(propertySources ...feather_commons_properties.PropertySource)
}

//

type EnvVar string

func NewEnvVar(value string) EnvVar {
	return EnvVar(value)
}

func (envVar EnvVar) AsInt() (int, error) {
	value, err := strconv.Atoi(string(envVar))
	if err != nil {
		return 0, err
	}
	return value, nil
}

func (envVar EnvVar) AsString() string {
	return string(envVar)
}
