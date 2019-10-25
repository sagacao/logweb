package db

import (
	"fmt"
	"logweb/ini"
	"logweb/models/base"
	"logweb/utils"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/henrylee2cn/faygo"
)

func GetOnlinedata(zeusid string, stm time.Time, results map[int]*base.OnlineData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}
	tablename := reader.gettablename("onlinestate", stm)

	sql := fmt.Sprintf("select UNIX_TIMESTAMP(logtime)-UNIX_TIMESTAMP(DATE(logtime)) as sec , online_user from %s; ", tablename)

	for i := 0; i < 24; i++ {
		for j := 0; j < 12; j++ {
			sec := j + 12*i
			results[sec] = &base.OnlineData{
				Time:    sec,
				RoleNum: 0,
			}
		}
	}

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("GetOnlinedata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	maxdata := 288
	for _, v := range buckets {
		sec := utils.GetInt("sec", v)

		faygo.Infof("sec %v ", sec)
		faygo.Infof("sec %v ", utils.GetInt("online_user", v))
		idx := sec / 300
		if idx >= 0 && idx < maxdata {
			results[idx].RoleNum = utils.GetInt("online_user", v)
		}
	}

	return nil
}

func getCurOnlineData(zeusid string, results map[string]*base.CurOnlineData) error {

	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	nowTime := time.Now()
	tablename := reader.gettablename("onlinestate", nowTime)

	sql := fmt.Sprintf("select online_user from %s order by logtime desc limit 1", tablename)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("GetOnlinedata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		num := utils.GetInt("online_user", v)
		results[zeusid].RoleNum = num
		break
	}

	return nil
}

func GetCurOnline(results map[string]*base.CurOnlineData) error {

	for k, v := range ini.ZeusType.Info {

		results[k] = &base.CurOnlineData{
			ServerID:   k,
			ServerName: v,
			RoleNum:    0,
		}
		getCurOnlineData(k, results)
	}

	return nil
}

func oneonlineingdata(engine *xorm.Engine, prof string, tm time.Time, results map[string]int) error {

	tablename := reader.gettablename("rolelogout", tm)

	sql := fmt.Sprintf("select ta.plat, ROUND(avg(ta.duration)) as data from %s as ta", tablename)
	sql += fmt.Sprintf(" where not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	if prof != "0" {
		sql += fmt.Sprintf(" and ta.profession='%s'", prof)
	}
	sql += fmt.Sprintf(" group by ta.plat;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("GetOnlinedata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {

		plat := utils.GetString("plat", v)
		_, ok := results[plat]
		if ok {
			results[plat] = utils.GetInt("data", v) / 60
		}
	}

	return nil
}

// func GetAvgOnlineData(zeusid string, prof string, stm, etm time.Time, results map[string]*base.AvgOnlineData) error {
func GetAvgOnlineData(zeusid string, prof string, stm, etm time.Time, results *[]*base.AvgOnlineData) error {

	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	for stm.Before(etm) {

		onlineData := make(map[string]int)
		_date := stm.Format("2006-01-02")

		for k, _ := range ini.Platform.Info {
			onlineData[k] = 0
		}

		oneonlineingdata(engine, prof, stm, onlineData)

		for k, v := range onlineData {
			/* results[_date] = &base.AvgOnlineData{
				Plat:      k,
				Date:      _date,
				OnlineMin: v,
			} */
			*results = append(*results, &base.AvgOnlineData{
				Plat:      k,
				Date:      _date,
				OnlineMin: v,
			})
		}
		stm = stm.Add(time.Hour * 24)
	}

	return nil
}

func GetLeagueLineData(zeusid string, stm time.Time, results map[int]*base.LeagueLineData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}
	tablename := reader.gettablename("onlinestate", stm)

	sql := fmt.Sprintf("select UNIX_TIMESTAMP(logtime)-UNIX_TIMESTAMP(DATE(logtime)) as sec , league_num from %s;", tablename)

	for i := 0; i < 24; i++ {
		for j := 0; j < 12; j++ {
			sec := j + 12*i
			results[sec] = &base.LeagueLineData{
				Time:      sec,
				LeagueNum: 0,
			}
		}
	}

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("GetOnlinedata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	maxdata := 288
	for _, v := range buckets {
		sec := utils.GetInt("sec", v)
		idx := sec / 300
		if idx >= 0 && idx < maxdata {
			results[idx].LeagueNum = utils.GetInt("league_num", v)
		}
	}

	return nil
}
