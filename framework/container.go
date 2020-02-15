package framework

import (
	"sync"

	"github.com/pkg/errors"
)

// Container is a core struct which store provider and instance
type Container interface {
	Bind(ServiceProvider, bool) error
	Singleton(ServiceProvider) error

	Make(string) (interface{}, error)
	MakeNew(string, []interface{}) (interface{}, error)
}

// HadeContainer is instance of Container
type HadeContainer struct {
	Container
	providers    []ServiceProvider
	instances    map[string]interface{}
	methods      map[string]NewInstance
	isSingletons map[string]bool

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

func (hade *HadeContainer) Bind(provider ServiceProvider, isSingleton bool) error {
	hade.lock.RLock()
	defer hade.lock.RUnlock()
	key := provider.Name()

	hade.providers = append(hade.providers, provider)
	hade.isSingletons[key] = isSingleton
	hade.methods[key] = provider.Register(hade)

	// if provider is not defer
	if provider.IsDefer() == false {
		params := provider.Params()
		method := hade.methods[key]
		provider.Boot(hade)
		instance, err := method(params...)
		if err != nil {
			return errors.New(err.Error())
		}
		if isSingleton == true {
			hade.instances[key] = instance
		}
	}
	return nil
}

func (hade *HadeContainer) Singleton(provider ServiceProvider) error {
	return hade.Bind(provider, true)
}

func (hade *HadeContainer) FindServiceProvider(key string) ServiceProvider {
	for _, sp := range hade.providers {
		if sp.Name() == key {
			return sp
		}
	}
	return nil
}

func (hade *HadeContainer) Make(key string) (interface{}, error) {
	return hade.make(key, nil)
}

func (hade *HadeContainer) MakeNew(key string, params []interface{}) (interface{}, error) {
	return hade.make(key, params)
}

func (hade *HadeContainer) make(key string, params []interface{}) (interface{}, error) {
	// check has Register
	if hade.FindServiceProvider(key) == nil {
		return nil, nil
	}

	// check instance
	if ins, ok := hade.instances[key]; ok {
		return ins, nil
	}

	// is not instance
	method := hade.methods[key] // must ok
	prov := hade.FindServiceProvider(key)
	isSingle := hade.isSingletons[key]
	if params == nil {
		params = prov.Params()
	}
	prov.Boot(hade)
	ins, err := method(params...)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	if isSingle {
		hade.instances[key] = ins
		return ins, nil
	}
	return ins, nil
}
