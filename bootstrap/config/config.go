package config

import(
	"github.com/spf13/viper"
)

// Default 是默认的配置
var Default *viper.Viper

// Init 初始化加载Config
func Init(config string) error {
	Default = viper.New()
	Default.SetConfigType("yaml")
	Default.SetConfigFile(config)
	if err := Default.ReadInConfig(); err != nil {
		return err	
	}
	return nil
}