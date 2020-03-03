package gorm

import (
	"testing"

	"github.com/jianfengye/hade/framework"
	"github.com/jianfengye/hade/framework/contract"
	"github.com/jianfengye/hade/framework/provider/app"
	"github.com/jianfengye/hade/framework/provider/config"
	"github.com/jianfengye/hade/framework/provider/env"
	"github.com/jianfengye/hade/framework/provider/log"
	"github.com/jianfengye/hade/tests"
	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func TestGorm_Normal(t *testing.T) {
	Convey("test gorm normal case", t, func() {
		basePath := tests.BasePath
		c := framework.NewHadeContainer()
		c.Singleton(&app.HadeAppProvider{BasePath: basePath})
		c.Singleton(&env.HadeEnvProvider{})
		c.Singleton(&config.HadeConfigProvider{})
		c.Singleton(&log.HadeLogServiceProvider{})

		err := c.Bind(&GormServiceProvider{}, false)
		So(err, ShouldBeNil)
		conf := c.MustMake(contract.ConfigKey).(contract.Config)
		// read default
		gm, err := c.MakeNew("gorm", []interface{}{conf.GetStringMapString("database.default")})
		So(err, ShouldBeNil)
		db := gm.(*gorm.DB)
		db.AutoMigrate(&Product{})
		db.Create(&Product{Code: "L1212", Price: 1000})
		var product Product
		db.First(&product, 1)                   // 查询id为1的product
		db.First(&product, "code = ?", "L1212") // 查询code为l1212的product
		db.Model(&product).Update("Price", 2000)
		db.Delete(&product)
		db.DropTableIfExists(&Product{})
	})
}
