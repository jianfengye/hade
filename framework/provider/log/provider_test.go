package log

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/jianfengye/hade/framework"
	"github.com/jianfengye/hade/framework/contract"
	"github.com/jianfengye/hade/framework/provider/app"
	"github.com/jianfengye/hade/framework/provider/config"
	"github.com/jianfengye/hade/framework/provider/env"
	"github.com/jianfengye/hade/framework/provider/log/formatter"
	"github.com/jianfengye/hade/tests"
	. "github.com/smartystreets/goconvey/convey"
	"golang.org/x/net/context"
)

func TestConsoleLog_Normal(t *testing.T) {
	Convey("test hade console log normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewHadeContainer()
		c.Singleton(&app.HadeAppProvider{BasePath: basePath})
		c.Singleton(&env.HadeEnvProvider{})
		c.Singleton(&config.HadeConfigProvider{})

		err := c.Singleton(&HadeLogServiceProvider{})
		So(err, ShouldBeNil)

		var buf bytes.Buffer
		l := c.MustMake(contract.LogKey).(contract.ConsoleLog)
		l.SetOutput(&buf)
		l.Debug(context.Background(), "test_debug", []interface{}{"foo1", []int{1, 2, 3}})
		l.Info(context.Background(), "test_debug", []interface{}{"foo2", []int{1, 2, 3}})
		So(buf.String(), ShouldNotContainSubstring, "foo1")
		So(buf.String(), ShouldContainSubstring, "foo2")

		buf.Reset()
		l.SetLevel(contract.InfoLevel)
		l.Debug(context.Background(), "test_debug", []interface{}{"foo1", []int{1, 2, 3}})
		l.Info(context.Background(), "test_debug", []interface{}{"foo2", []int{1, 2, 3}})
		So(buf.String(), ShouldNotContainSubstring, "foo1")

		buf.Reset()
		l.SetLevel(contract.InfoLevel)
		l.SetFormatter(formatter.JsonFormatter)
		ck := "foo"
		cv := "bar"
		ctx := context.WithValue(context.Background(), ck, cv)
		l.SetCxtFielder(func(ctx context.Context) []interface{} {
			v := ctx.Value(ck)
			return []interface{}{v}
		})
		l.Info(ctx, "test_console", []interface{}{"foo", []int{1, 2, 3}})
		So(buf.String(), ShouldContainSubstring, "[\"foo")
		So(buf.String(), ShouldContainSubstring, "bar")
	})
}

func TestSingleLog_Normal(t *testing.T) {
	Convey("test hade single log normal case", t, func() {
		basePath := tests.BasePath
		file := "hade_normal.log"

		c := framework.NewHadeContainer()
		c.Singleton(&app.HadeAppProvider{BasePath: basePath})
		c.Singleton(&env.HadeEnvProvider{})
		c.Singleton(&config.FakeConfigProvider{
			FileName: "log",
			Content:  []byte("driver: single\nfile: " + file),
		})
		app := c.MustMake(contract.AppKey).(contract.App)
		folder := app.LogPath()

		err := c.Singleton(&HadeLogServiceProvider{})
		So(err, ShouldBeNil)

		l := c.MustMake(contract.LogKey).(contract.SingleFileLog)
		// check file exist first
		l.Info(context.Background(), "test_single", []interface{}{"foo"})
		f := filepath.Join(folder, file)
		defer os.Remove(f)
		fd, err := os.Stat(f)
		So(err, ShouldBeNil)
		So(fd.Size(), ShouldBeGreaterThan, 0)
	})
}

func TestRotateLog_Normal(t *testing.T) {
	Convey("test hade rotate log normal case", t, func() {
		basePath := tests.BasePath
		file := "hade_normal.log"

		c := framework.NewHadeContainer()
		c.Singleton(&app.HadeAppProvider{BasePath: basePath})
		c.Singleton(&env.HadeEnvProvider{})
		c.Singleton(&config.FakeConfigProvider{
			FileName: "log",
			Content:  []byte("driver: rotate\nfile: " + file + "\nmax_files: 2\ndate_format: \"%Y%m%d\""),
		})
		app := c.MustMake(contract.AppKey).(contract.App)
		folder := app.LogPath()

		err := c.Singleton(&HadeLogServiceProvider{})
		So(err, ShouldBeNil)

		l := c.MustMake(contract.LogKey).(contract.RotatingFileLog)
		// check file exist first
		l.Info(context.Background(), "test_rotate", []interface{}{"foo"})
		f := filepath.Join(folder, file)
		f2 := filepath.Join(folder, file+"."+time.Now().Format("20060102"))
		defer os.Remove(f)
		defer os.Remove(f2)
		_, err = os.Stat(f)
		So(err, ShouldBeNil)
		fd2, err := os.Stat(f2)
		So(err, ShouldBeNil)
		So(fd2.Size(), ShouldBeGreaterThan, 0)
	})
}
