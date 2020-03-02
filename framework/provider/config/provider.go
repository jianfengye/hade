package config

import (
	"github.com/jianfengye/hade/framework"
	"github.com/jianfengye/hade/framework/contract"
)

type HadeConfigProvider struct {
	Folder string

	envMaps map[string]string
}

// Register registe a new function for make a service instance
func (provider *HadeConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeConfig
}

// Boot will called when the service instantiate
func (provider *HadeConfigProvider) Boot(c framework.Container) {
	if provider.Folder == "" && c.IsBind(contract.AppKey) {
		provider.Folder = c.MustMake(contract.AppKey).(contract.App).ConfigPath()
	}
	if c.IsBind(contract.EnvKey) {
		provider.envMaps = c.MustMake(contract.EnvKey).(contract.Env).All()
	}
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HadeConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *HadeConfigProvider) Params() []interface{} {
	return []interface{}{provider.Folder, provider.envMaps}
}

/// Name define the name for this service
func (provider *HadeConfigProvider) Name() string {
	return contract.ConfigKey
}
