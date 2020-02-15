package contract

// 配置接口
type Confige interface {
	// 根据string获取配置
	Get(string) interface{}
	// 判断是否存在
	IsExist(string) bool

	// 获取int，无则返回0
	GetInt(string) int
	// 获取string，无则返回空字符串
	GetStr(string) string
}
