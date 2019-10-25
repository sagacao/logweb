package mongo

import (
	. "logweb/ini"
	"logweb/utils"
	"sync"

	"github.com/henrylee2cn/faygo"
	"gopkg.in/mgo.v2"
)

type MongoEngine struct {
	lock   sync.Mutex
	islog  bool
	config MongoConfig

	mgos       utils.Pool
	mgoSession *mgo.Session
}

var mongoengine *MongoEngine

func (me *MongoEngine) init() {
	me.config = Config.Mongo

	if me.config.ShowLog == 1 {
		me.showLog(true)
	}

	var err error
	me.mgoSession, err = mgo.Dial(me.config.URL)
	if err != nil {
		faygo.Criticalf("init mongo connect error (%v)(%v)", me.config.URL, err)
		panic(err)
	}
}

func (me *MongoEngine) showLog(show bool) {
	me.islog = show
}

func (me *MongoEngine) getSession() *mgo.Session {
	if me.mgoSession == nil {
		var err error
		me.mgoSession, err = mgo.Dial(me.config.URL)
		if err != nil {
			return nil
		}
	}
	return me.mgoSession.Clone()
}

func (me *MongoEngine) witchCollection(dataBase, collection string, s func(*mgo.Collection) error) error {
	session := me.getSession()
	defer session.Close()
	c := session.DB(dataBase).C(collection)
	return s(c)
}

func Run() {
	mongoengine = &MongoEngine{}
	mongoengine.init()
}
