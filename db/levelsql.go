package db

import (
	"fmt"
	"logweb/models/level"
	"logweb/utils"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/henrylee2cn/faygo"
)

func getnewlyleveldisdata(engine *xorm.Engine, stm, etm time.Time, prof int, series *[]*level.LevelDisData) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "roleinfo"
	sql := fmt.Sprintf("select ta.level, count(ta.level) as num from %s as ta", tablename)
	sql += fmt.Sprintf(" where not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	sql += fmt.Sprintf(" and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	if prof != 0 {
		sql += fmt.Sprintf(" and ta.profession = '%v'", prof)
	}
	sql += fmt.Sprintf(" group by ta.level;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getnewlyplatdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}

	for _, v := range buckets {
		*series = append(*series, &level.LevelDisData{
			Plat:  0,
			Level: utils.GetInt64("level", v),
			Num:   utils.GetInt64("num", v),
		})
	}

	return nil

}

func getnewlylevellosedata(engine *xorm.Engine, stm, etm time.Time, prof int, series *[]*level.LevelLoseData) error {

	end_day := stm.Add(24 * time.Hour)
	startday := stm.Format("2006-01-02")
	cteat_role_end_day := end_day.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "roleinfo"
	sql := fmt.Sprintf("select ta.plat, ta.level, count(ta.level) as num from %s as ta", tablename)
	sql += fmt.Sprintf(" where not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	sql += fmt.Sprintf(" and DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') < '%s' ", startday, cteat_role_end_day)
	sql += fmt.Sprintf(" and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", endday)
	if prof != 0 {
		sql += fmt.Sprintf(" and ta.profession = '%v'", prof)
	}
	sql += fmt.Sprintf(" group by ta.plat, ta.level;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getnewlyplatdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		*series = append(*series, &level.LevelLoseData{
			Plat:  utils.GetInt("plat", v),
			Level: utils.GetInt64("level", v),
			Num:   utils.GetInt64("num", v),
		})
	}

	return nil
}
