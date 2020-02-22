package app

import (
	"testing"

	"github.com/jianfengye/hade/framework"
	"github.com/jianfengye/hade/framework/contract"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeAppProvider(t *testing.T) {
	Convey("test normal case", t, func() {
		c := framework.NewHadeContainer()
		sp := &HadeAppProvider{}

		err := c.Singleton(sp)
		So(err, ShouldBeNil)

		app, err := c.Make("app")
		So(err, ShouldBeNil)
		var iapp *contract.App
		So(app, ShouldImplement, iapp)
		hade := app.(contract.App)
		logPath := hade.LogPath()
		So(logPath, ShouldEqual, "/Users/didi/Documents/workspace/hade/storage/logs/")
	})
}