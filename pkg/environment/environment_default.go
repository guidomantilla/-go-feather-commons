package environment

import feather_commons_properties "github.com/guidomantilla/go-feather-commons/pkg/properties"

const (
	OsPropertySourceName  = "OS_PROPERTY_SOURCE_NAME"
	CmdPropertySourceName = "CMD_PROPERTY_SOURCE_NAME" //nolint:gosec
)

type DefaultEnvironmentOption func(environment *DefaultEnvironment)

func WithArrays(osArgsArray *[]string, cmdArgsArray *[]string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		osSource := feather_commons_properties.NewDefaultPropertySource(OsPropertySourceName, feather_commons_properties.NewDefaultProperties(feather_commons_properties.FromArray(osArgsArray)))
		cmdSource := feather_commons_properties.NewDefaultPropertySource(CmdPropertySourceName, feather_commons_properties.NewDefaultProperties(feather_commons_properties.FromArray(cmdArgsArray)))
		environment.propertySources = append(environment.propertySources, osSource, cmdSource)
	}
}

func WithArraySource(name string, array *[]string) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		source := feather_commons_properties.NewDefaultPropertySource(name, feather_commons_properties.NewDefaultProperties(feather_commons_properties.FromArray(array)))
		environment.propertySources = append(environment.propertySources, source)
	}
}

func WithPropertySources(propertySources ...feather_commons_properties.PropertySource) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		environment.propertySources = propertySources
	}
}

type DefaultEnvironment struct {
	propertySources []feather_commons_properties.PropertySource
}

func NewDefaultEnvironment(options ...DefaultEnvironmentOption) *DefaultEnvironment {
	environment := &DefaultEnvironment{
		propertySources: make([]feather_commons_properties.PropertySource, 0),
	}
	for _, opt := range options {
		opt(environment)
	}

	return environment
}

func (environment *DefaultEnvironment) GetValue(property string) EnvVar {

	var value string
	for _, source := range environment.propertySources {
		internalValue := source.Get(property)
		if internalValue != "" {
			value = internalValue
			break
		}
	}
	return NewEnvVar(value)
}

func (environment *DefaultEnvironment) GetValueOrDefault(property string, defaultValue string) EnvVar {

	envVar := environment.GetValue(property)
	if envVar != "" {
		return envVar
	}
	return NewEnvVar(defaultValue)
}

func (environment *DefaultEnvironment) GetPropertySources() []feather_commons_properties.PropertySource {
	return environment.propertySources
}

func (environment *DefaultEnvironment) AppendPropertySources(propertySources ...feather_commons_properties.PropertySource) {
	environment.propertySources = append(environment.propertySources, propertySources...)
}
