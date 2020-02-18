pacakge framework

func TestHadeContainer_normal(t *testing.T) {
	var err error
	c := NewHadeContainer()

	sp := &demo.DemoServiceProvider{
		C: map[string]string {"foo" : "bar"}
	}
	err = c.Bind(sp, true)
	if err != nil {
		
	}	
}