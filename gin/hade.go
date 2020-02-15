package gin

import "github.com/jianfengye/hade/framework"

// Hade framework add functions

// Register register a service provider for hade framework
func Register(engine *Engine, provider framework.ServiceProvider, isSingleton bool) {
	engine.container.Bind(provider, true)
}
