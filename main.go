package main

import (
	"logweb/db"
	_ "logweb/ini"
	"logweb/mongo"
	"logweb/router"

	"github.com/henrylee2cn/faygo"
)

func main() {
	db.Run()
	//elastic.Run()
	mongo.Run()

	router.Route(faygo.New("logweb"))
	faygo.Run()
}
