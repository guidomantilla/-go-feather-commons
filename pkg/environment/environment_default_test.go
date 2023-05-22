package environment

import (
	"testing"

	"github.com/stretchr/testify/assert"

	feather_commons_properties "github.com/guidomantilla/go-feather-commons/pkg/properties"
)

func Test_NewDefaultEnvironment(t *testing.T) {

	source := feather_commons_properties.NewDefaultPropertySource("some_property_source", feather_commons_properties.NewDefaultProperties())
	environment := NewDefaultEnvironment(WithPropertySources(source))

	assert.NotNil(t, environment)
	assert.NotNil(t, environment.propertySources)
	assert.NotEmpty(t, environment.propertySources)
	assert.Equal(t, "some_property_source", environment.propertySources[0].AsMap()["name"])
}

func Test_GetValue(t *testing.T) {

	props := feather_commons_properties.NewDefaultProperties()
	props.Add("some_property", "some_value")

	source := feather_commons_properties.NewDefaultPropertySource("some_property_source", props)
	environment := NewDefaultEnvironment(WithPropertySources(source))

	value := environment.GetValue("some_property")
	assert.NotNil(t, value)
	assert.NotEmpty(t, value)
	assert.Equal(t, "some_value", value.AsString())
}

func Test_GetValueOrDefault_ReturnDefault(t *testing.T) {

	props := feather_commons_properties.NewDefaultProperties()

	source := feather_commons_properties.NewDefaultPropertySource("some_property_source", props)
	environment := NewDefaultEnvironment(WithPropertySources(source))

	value := environment.GetValueOrDefault("some_property", "some_default_value")
	assert.NotNil(t, value)
	assert.NotEmpty(t, value)
	assert.Equal(t, "some_default_value", value.AsString())
}

func Test_GetValueOrDefault_ReturnValue(t *testing.T) {

	props := feather_commons_properties.NewDefaultProperties()
	props.Add("some_property", "some_value")

	source := feather_commons_properties.NewDefaultPropertySource("some_property_source", props)
	environment := NewDefaultEnvironment(WithPropertySources(source))

	value := environment.GetValueOrDefault("some_property", "some_default_value")
	assert.NotNil(t, value)
	assert.NotEmpty(t, value)
	assert.Equal(t, "some_value", value.AsString())
}

func Test_GetPropertySources(t *testing.T) {

	props := feather_commons_properties.NewDefaultProperties()
	props.Add("some_property", "some_value")

	source := feather_commons_properties.NewDefaultPropertySource("some_property_source", props)
	environment := NewDefaultEnvironment(WithPropertySources(source))

	propertySources := environment.GetPropertySources()
	assert.NotNil(t, propertySources)
	assert.NotEmpty(t, propertySources)
	assert.Equal(t, "some_property_source", environment.propertySources[0].AsMap()["name"])
}

func Test_AppendPropertySources(t *testing.T) {

	props1 := feather_commons_properties.NewDefaultProperties()
	props1.Add("some_property", "some_value")

	source1 := feather_commons_properties.NewDefaultPropertySource("some_property_source1", props1)
	environment := NewDefaultEnvironment(WithPropertySources(source1))

	props2 := feather_commons_properties.NewDefaultProperties()
	source2 := feather_commons_properties.NewDefaultPropertySource("some_property_source2", props2)

	environment.AppendPropertySources(source2)

	propertySources := environment.GetPropertySources()

	assert.NotNil(t, propertySources)
	assert.NotEmpty(t, propertySources)
	assert.Equal(t, "some_property_source1", environment.propertySources[0].AsMap()["name"])
	assert.Equal(t, "some_property_source2", environment.propertySources[1].AsMap()["name"])
}
