package environment

import feather_commons_properties "github.com/guidomantilla/go-feather-commons/pkg/properties"

// DefaultEnvironment

type DefaultEnvironment struct {
	propertySources []feather_commons_properties.PropertySource
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

//

type DefaultEnvironmentOption = func(environment *DefaultEnvironment)

func NewDefaultEnvironment(options ...DefaultEnvironmentOption) *DefaultEnvironment {
	environment := &DefaultEnvironment{
		propertySources: make([]feather_commons_properties.PropertySource, 0),
	}
	for _, opt := range options {
		opt(environment)
	}

	return environment
}

func WithPropertySources(propertySources ...feather_commons_properties.PropertySource) DefaultEnvironmentOption {
	return func(environment *DefaultEnvironment) {
		environment.propertySources = propertySources
	}
}
