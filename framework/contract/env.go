package contract

const (
	EnvProduction  = "production"
	EnvTesting     = "testing"
	EnvDevelopment = "development"
)

// Env define golang run enviornment
// it set some config which want ignored in git
type Env interface {
	AppEnv() string  // get current environment
	AppName() string // app name
	AppDebug() bool  // check app is debug open
	AppURL() string  // app url in local

	IsExist(string) bool // check environment setting exist
	Get(string) string   // get environment setting, if not exist, return ""
}
