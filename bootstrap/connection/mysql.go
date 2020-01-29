package connection

import (
	"github.com/spf13/viper"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"fmt"
)

// Default 获取默认的DB
var Default *gorm.DB

// Init 初始化数据库
func Init(vconf *viper.Viper) error {
	username := vconf.GetString("mysql.username")
	password := vconf.GetString("mysql.password")
	host := vconf.GetString("mysql.host")
	port := vconf.GetString("mysql.port")
	database := vconf.GetString("mysql.database")
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, database))
	if err != nil {
		return err
	}

	db.LogMode(true)
	Default = db
	return nil
}

// Refresh 更新数据库
func Refresh(vconf *viper.Viper) error {
	Destory(vconf)
	return Init(vconf)
}

// Destory 关掉数据库
func Destory(vconf *viper.Viper) {
	Default.Close()
}