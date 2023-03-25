package tools

//go:generate mockgen -source ../pkg/environment/types.go -destination ../pkg/environment/mocks.go -package=environment
//go:generate mockgen -source ../pkg/properties/types.go -destination ../pkg/properties/mocks.go -package=properties
//go:generate mockgen -source ../pkg/security/types.go -destination ../pkg/security/mocks.go -package=security
