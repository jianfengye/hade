package framework

import (
	"github.com/jianfengye/hade/framework/contract"
)

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
