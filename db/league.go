package db

import (
	"fmt"
	"logweb/models/league"
	"time"

	"github.com/henrylee2cn/faygo"
)

func GetLeagueData(zeusid string, stm, etm time.Time, results *[]*league.LeagueData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}
	
	getleaguedata(engine, results)
	return nil
	
}

func GetLeagueInfo(zeusid string, stm time.Time, number int, results *[]*league.LeagueInfo) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}
	
	getleagueinfo(engine, stm, number, results)
	return nil
}