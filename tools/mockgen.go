package tools

//go:generate mockgen -source ../pkg/environment/types.go -destination ../pkg/environment/mocks.go -package=environment
//go:generate mockgen -source ../pkg/properties/types.go -destination ../pkg/properties/mocks.go -package=properties
//go:generate mockgen -source ../pkg/log/types.go -destination ../pkg/log/mocks.go -package=log
//go:generate mockgen -source ../pkg/resilience/types.go -destination ../pkg/resilience/mocks.go -package=resilience
//go:generate mockgen -package config -destination ../pkg/config/mocks.go github.com/sethvargo/go-envconfig Lookuper
