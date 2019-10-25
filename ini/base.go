package ini

import (
	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/ini"
)

type (
	WebConfig struct {
		Mysql   MysqlConfig   `ini:"mysql" comment:"mysql security section"`
		Elastic ElasticConfig `ini:"elastic" comment:"elastic security section"`
		Mongo   MongoConfig   `ini:"mongo" comment:"mongo security section"`
	}

	MysqlConfig struct {
		Path     string `ini:"path" comment:"mysql path"`
		DataBase string `ini:"database" comment:"mysql database"`
		Host     string `ini:"host" comment:"mysql host"`
		Port     string `ini:"port" comment:"mysql port"`
		User     string `ini:"user" comment:"mysql user"`
		Password string `ini:"password" comment:"mysql password"`
		ShowLog  uint32 `ini:"showlog" comment:"mysql showlog"`
	}

	ElasticConfig struct {
		Host    string `ini:"host" comment:"elastic host"`
		Port    string `ini:"port" comment:"elastic port"`
		UsePool uint32 `ini:"usepool" comment:"elastic usepool"`
		ShowLog uint32 `ini:"showlog" comment:"elastic showlog"`
	}

	MongoConfig struct {
		URL     string `ini:"url" comment:"mongo host"`
		ShowLog uint32 `ini:"showlog" comment:"mongo showlog"`
	}
)

var (
	Config WebConfig
)

func init() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		panic("config is nil")
	}

	err = cfg.Section("mysql").MapTo(&Config.Mysql)
	if err != nil {
		panic("mysql config is error")
	}
	faygo.Infof("init mysql : %s, %s", Config.Mysql.Host, Config.Mysql.Port)

	err = cfg.Section("elastic").MapTo(&Config.Elastic)
	if err != nil {
		panic("elastic config is error")
	}
	faygo.Infof("init elastic : %s, %s", Config.Elastic.Host, Config.Elastic.Port)

	err = cfg.Section("mongo").MapTo(&Config.Mongo)
	if err != nil {
		panic("mongo config is error")
	}
	faygo.Infof("init mongo : %s", Config.Mongo.URL)

	initZeusType()
	initPlatform()
	initItemChange()
	initRechargeform()
}
