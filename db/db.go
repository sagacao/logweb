package db

import (
	"fmt"
	. "logweb/ini"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type dbReader struct {
	islog  bool
	config MysqlConfig

	reader *utils.Map
}

var reader *dbReader

func (db *dbReader) getdbsource(dbname string) string {
	//dbconn := db.config.User + ":" + db.config.Password + "@tcp(" + db.config.Host + ")/" + dbname + "?charset=utf8"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		db.config.User, db.config.Password, db.config.Host, db.config.Port, dbname)
}

func (db *dbReader) gettablename(tablename string, tm time.Time) string {
	return fmt.Sprintf("log_%s_%s", tm.Format("20060102"), tablename)
}

func (db *dbReader) init() {
	db.config = Config.Mysql

	if db.config.ShowLog == 1 {
		db.showLog(true)
	}
}

func (db *dbReader) getReader(zone string) *xorm.Engine {
	readerif := db.reader.Get(zone)
	if readerif == nil {
		return db.produceReader(zone)
	}

	reader, ok := readerif.(*xorm.Engine)
	if ok && reader != nil {
		return reader
	}
	db.reader.Del(zone)
	return db.produceReader(zone)
}

func (db *dbReader) produceReader(zone string) *xorm.Engine {
	zoneid := utils.ParseUint32(zone)
	dbname := fmt.Sprintf("%s_%v", Config.Mysql.DataBase, zoneid) //getdbname(zoneid)

	dataSource := db.getdbsource(dbname)
	engine, err := xorm.NewEngine("mysql", dataSource)
	if err != nil {
		faygo.Warningf("produce db engine '%s' Error : (%v)", dbname, err)
		return nil
	}

	db.reader.Set(zone, engine)
	return engine
}

func (db *dbReader) showLog(show bool) {
	reader.islog = show
}

func printsql(sqlinfo interface{}) {
	if reader.islog {
		faygo.Infof("sql:`%v`", sqlinfo)
	}
}

func Run() {
	reader = &dbReader{islog: false, reader: new(utils.Map)}
	reader.init()
}
