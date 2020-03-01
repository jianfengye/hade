package config

import (
	"github.com/jianfengye/hade/framework"
	"github.com/jianfengye/hade/framework/contract"
)

type FakeConfigProvider struct {
	FileName string
	Content  []byte
}

// Register registe a new function for make a service instance
func (provider *FakeConfigProvider) Register(c framework.Container) framework.NewInstance {
	return NewFakeConfig
}

// Boot will called when the service instantiate
func (provider *FakeConfigProvider) Boot(c framework.Container) {
}

// IsDefer define whether the service instantiate when first make or register
func (provider *FakeConfigProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *FakeConfigProvider) Params() []interface{} {
	return []interface{}{provider.FileName, provider.Content}
}

// Name define the name for this service
func (provider *FakeConfigProvider) Name() string {
	return contract.ConfigKey
}
