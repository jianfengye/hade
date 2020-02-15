package gin

import "github.com/jianfengye/hade/framework"

// Hade framework add functions

// Register register a service provider for hade framework
func Register(engine *Engine, provider framework.ServiceProvider, isSingleton bool) {
	engine.container.Bind(provider, isSingleton)
}

// Register register a singleton serviceProvider
func RegisterSingleton(engine *Engine, provider framework.ServiceProvider) {
	engine.container.Bind(provider, true)
}
