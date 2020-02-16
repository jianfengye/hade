package framework

import (
	"github.com/jianfengye/hade/framework/contract"
)

// ConfigService get ConfigService which provider by name
func (hade *HadeContainer) ConfigService() (contract.Config, error) {
	ins, err := hade.Make("config")
	if err != nil {
		return nil, err
	}
	if ins != nil {
		if c, ok := ins.(contract.Config); ok {
			return c, nil
		}
	}
	return nil, nil
}
