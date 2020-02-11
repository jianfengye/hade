package framework

import (
	"github.com/jianfengye/hade/framework/contract"
	serviceProvider "github.com/jianfengye/hade/framework/provider"
)

const (
	CONFIG_SERVICE: "config"
)

type Provider func(...interface{}) (interface{})

// 注入器
type Injector interface {
	SetProvider(string, Provider)

	ServiceConfig(...interface{}) contract.IConfiger
}

type HadeInjector struct {
	instances map[string]interface{} // 单例
	providers map[string]Provider
}

func NewHadeInjector() *HadeInjector {
	return &HadeInjector{
		instances: map[string]interace{},
		providers: map[string]Provider{} 
	}
}

func (hade *HadeInjector) SetProvider(key string, provider Provider) {
	hade.providers[key] = provider
	return
}

func (hade *HadeInjector) ServiceConfig(options ...interface{}) contract.IConfiger {
	key := CONFIG_SERVICE
	if v := hade.instances(key); v != nil {
		return v.(contract.IConfiger)
	}

	var fn Provider
	if existNewProvider, ok := hade.providers[key]; ok {
		fn = existNewProvider
	} else {
		fn = serviceProvider.NewViperProvider
	}

	c := fn(options).(contract.IConfiger)
	hades.providers[key] = fn
	hades.instances[key] = c
	return c
}
