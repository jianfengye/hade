package gin

import "github.com/jianfengye/hade/framework"

// Hade framework add functions

func Register(engine *Engine, provider framework.ServiceProvider, isSingleton bool) {
	engine.container.Bind(provider, true)
}
