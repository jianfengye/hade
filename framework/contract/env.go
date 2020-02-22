package contract

const (
	// EnvProduction represent the environment which build for production
	EnvProduction = "production"
	// EnvTesting represent the environment which build for test
	EnvTesting = "testing"
	// EnvDevelopment represent the environment which build for development
	EnvDevelopment = "development"

	// EnvKey is the key in container
	EnvKey = "env"
)

// Env define golang run enviornment
// it set some config which want ignored in git
type Env interface {
	// AppEnv get current environment
	AppEnv() string
	// AppDebug check app is debug open
	AppDebug() bool
	// AppURL define app url in local
	AppURL() string

	// IsExist check setting is exist
	IsExist(string) bool
	// Get environment setting, if not exist, return ""
	Get(string) string
}
