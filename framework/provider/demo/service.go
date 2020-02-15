package demo

type DemoService struct {
	c map[string]string
}

func NewDemoService(params ...interface{}) (interface{}, error) {
	c := params[0].(map[string]string)
	return &DemoService{c: c}, nil
}

func (s *DemoService) Get(key string) string {
	if v, ok := s.c[key]; ok {
		return v
	}
	return ""
}
