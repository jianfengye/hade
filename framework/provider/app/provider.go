package app

import (
	"github.com/jianfengye/hade/framework"
)

// HadeAppProvider provide a App service, it must be singlton, and not delay
type HadeAppProvider struct {
	app *HadeApp
}

// Register registe a new function for make a service instance
func (provider *HadeAppProvider) Register(c framework.Container) framework.NewInstance {
	return NewHadeApp
}

// Boot will called when the service instantiate
func (provider *HadeAppProvider) Boot(c framework.Container) {
}

// IsDefer define whether the service instantiate when first make or register
func (provider *HadeAppProvider) IsDefer() bool {
	return false
}

// Params define the necessary params for NewInstance
func (provider *HadeAppProvider) Params() []interface{} {
	return nil
}

/// Name define the name for this service
func (provider *HadeAppProvider) Name() string {
	return "app"
}