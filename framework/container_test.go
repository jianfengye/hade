package framework

import (
	"fmt"
	"testing"

	"github.com/jianfengye/hade/framework/contract"
	. "github.com/smartystreets/goconvey/convey"
)

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

type DemoServiceProvider struct {
	C     map[string]string
	Defer bool
	ServiceProvider
}

func (sp *DemoServiceProvider) Name() string {
	return "demo"
}

func (sp *DemoServiceProvider) Register(c Container) NewInstance {
	return NewDemoService
}

func (sp *DemoServiceProvider) IsDefer() bool {
	return sp.Defer
}

func (sp *DemoServiceProvider) Params() []interface{} {
	return []interface{}{sp.C}
}

func (sp *DemoServiceProvider) Boot(c Container) {
	fmt.Println("demo service boot")
}

func TestHadeContainer_Singleton_NoDefer(t *testing.T) {
	Convey("test normal case", t, func() {

		Convey("create a hade container", nil)
		c := NewHadeContainer()

		Convey("create a demo service provider", nil)
		sp := &DemoServiceProvider{
			C: map[string]string{"foo": "bar"},
		}
		err := c.Singleton(sp)
		Convey("register demo service provider to container", func() {
			ShouldBeNil(err)
		})

		Convey("make a demo service", func() {
			demo, err := c.Make("demo")
			ShouldBeNil(err)
			serv, ok := demo.(contract.Demo)
			ShouldBeTrue(ok)
			Convey("call demo service method get ok", func() {
				val := serv.Get("foo")
				ShouldEqual(val, "bar")
			})
		})
	})
}

func TestHadeContainer_Singleton_Defer(t *testing.T) {
	Convey("test normal case", t, func() {

		Convey("create a hade container", nil)
		c := NewHadeContainer()

		Convey("create a demo service provider", nil)
		sp := &DemoServiceProvider{
			C:     map[string]string{"foo": "bar"},
			Defer: true,
		}
		err := c.Singleton(sp)
		Convey("register demo service provider to container", func() {
			ShouldBeNil(err)
		})

		Convey("make a demo service", func() {
			demo, err := c.Make("demo")
			ShouldBeNil(err)
			serv, ok := demo.(contract.Demo)
			ShouldBeTrue(ok)
			Convey("call demo service method get ok", func() {
				val := serv.Get("foo")
				ShouldEqual(val, "bar")
			})
		})
	})
}
func TestHadeContainer_NotSinglton_Defer(t *testing.T) {
	Convey("test normal case", t, func() {

		Convey("create a hade container", nil)
		c := NewHadeContainer()

		Convey("create a demo service provider", nil)
		sp := &DemoServiceProvider{
			C:     map[string]string{"foo": "bar"},
			Defer: false,
		}
		err := c.Bind(sp, false)
		Convey("register demo service provider to container", func() {
			ShouldBeNil(err)
		})

		Convey("make a demo service", func() {
			demo, err := c.Make("demo")
			ShouldBeNil(err)
			serv, ok := demo.(contract.Demo)
			ShouldBeTrue(ok)
			Convey("call demo service method get ok", func() {
				val := serv.Get("foo")
				ShouldEqual(val, "bar")
			})
		})
	})
}
