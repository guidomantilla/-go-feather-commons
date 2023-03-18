package tools

//go:generate mockgen -package environment 	-destination ../pkg/environment/mocks.go	../pkg/environment Environment
//go:generate mockgen -package properties 	-destination ../pkg/properties/mocks.go 	../pkg/properties Properties,PropertySource
