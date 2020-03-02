package config

import (
	"path/filepath"
	"testing"

	"github.com/jianfengye/hade/tests"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHadeConfig_GetInt(t *testing.T) {
	Convey("test hade env normal case", t, func() {
		basePath := tests.BasePath
		folder := filepath.Join(basePath, "config")
		serv, err := NewHadeConfig(folder)
		So(err, ShouldBeNil)
		conf := serv.(*HadeConfig)
		timeout := conf.GetInt("database.mysql.timeout")
		So(timeout, ShouldEqual, 1)
	})
}
