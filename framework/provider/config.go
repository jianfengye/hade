package provider

import (
	"github.com/spf13/viper"
)

type ViperProvider struct {
	viper *viper.Viper
}

func DefaultViperProvider(options ...interface{}) interface{} {
	return nil
}

func NewViperProvider(options ...interface{}) interface{} {
	// TODO: generateViperProvider by options
	return nil
}

func (v *ViperProvider) Get(string) interface{} {
	return nil
}

func (v *ViperProvider) GetInt(string) int {
	return 0
}

func (v *ViperProvider) GetStr(string) string {
	return ""
}

func (v *ViperProvider) IsExist(string) bool {
	return false
}