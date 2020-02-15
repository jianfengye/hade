package framework

import (
	"sync"

	"github.com/jianfengye/hade/framework/contract"
)

var (
	CONFIG_SERVICE = "config"
)

// Provider all service provider should implement this for new service
type Provider func(...interface{}) interface{}

type providerKit struct {
	provider  Provider      // 注册器
	params    []interface{} // provider需要的参数
	service   interface{}   // 最终的服务
	isSington bool          // 判断是否单例子
}

// Container is a core struct which store provider and instance
type Container interface {
	SetProvider(string, Provider, bool, []interface{})
	GetService(string) interface{}
	NewService(string, []interface{})
}

// HadeContainer is instance of Container
type HadeContainer struct {
	kits map[string]providerKit // 测试套件
	lock sync.RWMutex
}

// NewHadeContainer is new instance
func NewHadeContainer() *HadeContainer {
	return &HadeContainer{
		kits: map[string]providerKit{},
		lock: sync.RWMutex{},
	}
}

// SetProvider set provider
func (hade *HadeContainer) SetProvider(key string, provider Provider, isSington bool, params []interface{}) {
	kit := providerKit{
		provider:  provider,
		params:    params,
		isSington: isSington,
	}
	hade.kits[key] = kit
	return
}

func (hade *HadeContainer) GetService(key string) interface{} {
	hade.lock.RLock()
	defer hade.lock.RUnlock()

	if k, ok := hade.kits[key]; ok {
		if !k.isSington {
			return k.provider(k.params...)
		} else {
			if k.service != nil {
				return k.service
			} else {
				k.service = k.provider(k.params...)
				return k.service
			}
		}
	}

	return nil
}

// ConfigService get ConfigService which provider by name
func (hade *HadeContainer) ConfigService() contract.Config {
	key := CONFIG_SERVICE
	ins := hade.GetService(key)
	if ins != nil {
		if c, ok := ins.(contract.Config); ok {
			return c
		}
	}
	return nil
}
