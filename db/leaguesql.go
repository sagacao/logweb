package db

import (
	"fmt"
	"logweb/models/league"
	"logweb/utils"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/henrylee2cn/faygo"
)

func getleaguedata(engine *xorm.Engine, series *[]*league.LeagueData) error {
	sql := fmt.Sprintf("select league, name, level, camp, leader, leader_id, state, brithday from leagueinfo order by level desc;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getleaguedata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		*series = append(*series, &league.LeagueData{
			Number:   utils.GetString("league", v),
			Level:    utils.GetString("level", v),
			Name:     utils.GetString("name", v),
			Camp:     utils.GetString("camp", v),
			Leader:   utils.GetString("leader", v),
			LeaderId: utils.GetString("leader_id", v),
			State:    utils.GetString("state", v),
			Brithday: utils.GetString("brithday", v),
		})
	}
	return nil

}

func getleagueinfo(engine *xorm.Engine, realtime time.Time, number int, series *[]*league.LeagueInfo) error {
	// timestr := realtime.Format("2006-01-02")

	tablename := reader.gettablename("leaguedata", realtime)
	sql := fmt.Sprintf("select logtime, level, assets, credits, health, member, student from %s where league=%d order by logtime;", tablename, number)
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getleaguedata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		*series = append(*series, &league.LeagueInfo{
			Logtime: utils.GetString("logtime", v),
			Level:   utils.GetInt64("level", v),
			Assets:  utils.GetInt64("assets", v),
			Credits: utils.GetInt64("credits", v),
			Health:  utils.GetInt64("health", v),
			Member:  utils.GetInt64("member", v),
			Student: utils.GetInt64("student", v),
		})
	}

	return nil
}
