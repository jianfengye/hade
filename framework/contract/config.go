package contract

import "time"

// Config define setting from files, it support key contains dov。
// for example:
// .Get("user.name")
// use toml-lang like, https://github.com/toml-lang/toml v0.4.0
type Config interface {
	IsExist(key string) bool

	Get(key string) interface{}
	GetBool(key string) bool
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetUint(key string) uint
	GetUint32(key string) uint32
	GetUint64(key string) uint64
	GetFloat64(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	GetIntSlice(key string) []int
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string

	Load(key string, val interface{}) error // load a config to a struct, val should be an pointer
}
