package framework

import (
	"sync"

	"github.com/jianfengye/hade/framework/contract"
)

type NewInstance func(...interface{}) (interface{}, error)

// Container is a core struct which store provider and instance
type Container interface {
	Bind(ServiceProvider, bool) Container
	Singleton(ServiceProvider) Container

	Make(string) interface{}
	MakeNew(string, []interface{}) interface{}

	ErrorChain
}

type ErrorChain interface {
	SetError(error)
	HasError() bool
	GetError() error
}

type DErrorChain struct {
	err error
}

func (e DErrorChain) SetError(err error) {
	e.err = err
}

func (e DErrorChain) HasError() bool {
	return e.err != nil
}

func (e DErrorChain) GetError() error {
	return e.err
}

type ServiceProvider interface {
	Register(Container) NewInstance
	Boot(Container)

	IsDefer() bool
	InstanceParams() []interface{}
	Name() string
}

// HadeContainer is instance of Container
type HadeContainer struct {
	Container
	providers    []ServiceProvider
	instances    map[string]interface{}
	methods      map[string]NewInstance
	isSingletons map[string]bool

	DErrorChain
	lock sync.RWMutex
}

// NewHadeContainer is new instance
func NewHadeContainer() *HadeContainer {
	return &HadeContainer{
		providers:    []ServiceProvider{},
		instances:    map[string]interface{}{},
		methods:      map[string]NewInstance{},
		isSingletons: map[string]bool{},
		lock:         sync.RWMutex{},
	}
}

func (hade *HadeContainer) GetError() error {
	return hade.DErrorChain.GetError()
}

func (hade *HadeContainer) HasError() bool {
	return hade.DErrorChain.HasError()
}

func (hade *HadeContainer) SetError(err error) {
	hade.SetError(err)
}

func (hade *HadeContainer) Bind(provider ServiceProvider, isSingleton bool) Container {
	hade.lock.RLock()
	defer hade.lock.RUnlock()
	key := provider.Name()

	hade.providers = append(hade.providers, provider)
	hade.isSingletons[key] = isSingleton
	hade.methods[key] = provider.Register(hade)

	// if provider is not defer
	if provider.IsDefer() == false {
		params := provider.InstanceParams()
		method := hade.methods[key]
		provider.Boot(hade)
		instance, err := method(params...)
		if err != nil {
			hade.SetError(err)
			return hade
		}
		if isSingleton == true {
			hade.instances[key] = instance
		}
	}
	return hade
}

func (hade *HadeContainer) Singleton(provider ServiceProvider) Container {
	hade.Bind(provider, true)
	return hade
}

func (hade *HadeContainer) FindServiceProvider(key string) ServiceProvider {
	for _, sp := range hade.providers {
		if sp.Name() == key {
			return sp
		}
	}
	return nil
}

func (hade *HadeContainer) Make(key string) interface{} {
	return hade.make(key, nil)
}

func (hade *HadeContainer) MakeNew(key string, params []interface{}) interface{} {
	return hade.make(key, params)
}

func (hade *HadeContainer) make(key string, params []interface{}) interface{} {
	// check has Register
	if hade.FindServiceProvider(key) == nil {
		return nil
	}

	// check instance
	if ins, ok := hade.instances[key]; ok {
		return ins
	}

	// is not instance
	method := hade.methods[key] // must ok
	prov := hade.FindServiceProvider(key)
	isSingle := hade.isSingletons[key]
	if params == nil {
		params = prov.InstanceParams()
	}
	prov.Boot(hade)
	ins, err := method(params...)
	if err != nil {
		hade.SetError(err)
		return nil
	}

	if isSingle {
		hade.instances[key] = ins
		return ins
	}
	return ins
}

// ConfigService get ConfigService which provider by name
func (hade *HadeContainer) ConfigService() contract.Config {
	ins := hade.Make("config")
	if ins != nil {
		if c, ok := ins.(contract.Config); ok {
			return c
		}
	}
	return nil
}
