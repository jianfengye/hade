package provider

type ViperProvider struct {
	params []interface{}
}

func (v *ViperProvider) Register(c Container) {
	c.BindMethod(v.Name(), NewViperService)
}

func (v *ViperProvider) Boot(c Container) {
}

func (v *ViperProvider) Params() []interface{} {
	return v.params
}

func (v *ViperProvider) Name() string {
	return "config"
}

type ViperService struct {
}

func NewViperService(...interface{}) (interface{}, error) {
	return nil, nil
}

func (v *ViperService) Get(string) interface{} {
	return nil
}

func (v *ViperService) GetInt(string) int {
	return 0
}

func (v *ViperService) GetStr(string) string {
	return ""
}

func (v *ViperService) IsExist(string) bool {
	return false
}
