package config

import (
	"os"

	"github.com/guidomantilla/go-feather-commons/pkg/environment"
	"github.com/guidomantilla/go-feather-commons/pkg/properties"
)

const (
	OsPropertySourceName  = "OS_PROPERTY_SOURCE_NAME"
	CmdPropertySourceName = "CMD_PROPERTY_SOURCE_NAME"
)

var (
	_environment environment.Environment
)

func Init(cmdArgs *[]string) environment.Environment {

	// Load CMD and OS variables
	osArgs := os.Environ()
	osSource := properties.NewDefaultPropertySource(OsPropertySourceName, properties.NewDefaultProperties(properties.FromArray(&osArgs)))
	cmdSource := properties.NewDefaultPropertySource(CmdPropertySourceName, properties.NewDefaultProperties(properties.FromArray(cmdArgs)))
	_environment = environment.NewDefaultEnvironment(environment.WithPropertySources(osSource, cmdSource))
	return _environment
}
