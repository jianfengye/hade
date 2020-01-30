package connection

import (
	"hade/bootstrap/logger"

	"github.com/spf13/viper"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DefaultMysql 获取默认的DB
var DefaultMysql *gorm.DB

// Init 初始化数据库
func Init(vconf *viper.Viper) error {

	// 如果有配置mysql，则设置连接池
	if vconf.Get("mysql") != nil {

		vconf.SetDefault("mysql.username", "hade")
		vconf.SetDefault("mysql.password", "hade")
		vconf.SetDefault("mysql.host", "127.0.0.1")
		vconf.SetDefault("mysql.port", "3306")

		username := vconf.GetString("mysql.username")
		password := vconf.GetString("mysql.password")
		host := vconf.GetString("mysql.host")
		port := vconf.GetString("mysql.port")

		vconf.SetDefault("mysql.database", "default")
		database := vconf.GetString("mysql.database")

		vconf.SetDefault("mysql.timeout", "5s")
		timeout := vconf.GetString("mysql.timeout")

		vconf.SetDefault("mysql.readTimeout", "10s")
		vconf.SetDefault("mysql.writeTimeout", "10s")
		readTimeout := vconf.GetString("mysql.readTimeout")
		writeTimeout := vconf.GetString("mysql.writeTimeout")

		vconf.SetDefault("mysql.maxOpenConns", 10)
		vconf.SetDefault("mysql.maxIdleConns", 10)
		maxOpenConns := vconf.GetInt("mysql.maxOpenConns")
		maxIdleConns := vconf.GetInt("mysql.maxIdleConns")

		db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&timeout=%s&readTimeout=%s&writeTimeout=%s", username, password, host, port, database, timeout, readTimeout, writeTimeout))
		if err == nil {
			db.SetLogger(logger.Default)
			db.DB().SetMaxOpenConns(maxOpenConns)
			db.DB().SetMaxIdleConns(maxIdleConns)

			db.LogMode(vconf.GetBool("mysql.logMode"))
			DefaultMysql = db
		} else {
			logger.Default.Error(err.Error())
			return err
		}

	}

	return nil
}

// Refresh 更新数据库
func Refresh(vconf *viper.Viper) error {
	Destory(vconf)
	return Init(vconf)
}

// Destory 关掉数据库
func Destory(vconf *viper.Viper) {
	if DefaultMysql != nil {
		DefaultMysql.Close()
	}
}
