package framework

type NewInstance func(...interface{}) (interface{}, error)

type ServiceProvider interface {
	Register(Container) NewInstance
	Boot(Container)

	IsDefer() bool
	Params() []interface{}
	Name() string
}
