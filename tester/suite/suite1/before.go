package suite1

import (
	"hade/bootstrap/config"
	"hade/bootstrap/connection"
	"hade/bootstrap/logger"
	"hade/model"
	"hade/util"
	"log"
	"path"
)

func SetUp() {
	cf := path.Join(util.RootFolder(), "env.test.yaml")
	if err := config.Init(cf); err != nil {
		log.Fatalln(err)
	}

	// 加载日志
	if err := logger.Init(config.Default); err != nil {
		log.Fatalln(err)
	}

	// 加载数据库
	if err := connection.Init(config.Default); err != nil {
		log.Fatalln(err)
	}
}


func Before() {
	// 清空所有数据
	connection.Default.Delete(&model.User{})
	connection.Default.Delete(&model.Topic{})
	connection.Default.Delete(&model.Tag{})
	connection.Default.Delete(&model.Comment{})
	connection.Default.Delete(&model.Likeable{})
	connection.Default.Delete(&model.Taggable{})
	initDB()
}
