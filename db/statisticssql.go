package db

import (
	"fmt"
	"logweb/models/base"
	"logweb/models/statistics"
	"logweb/utils"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/henrylee2cn/faygo"
)

func getnewlyplatdata(engine *xorm.Engine, stm time.Time, results map[string]*base.PlatDnuData) error {
	startday := stm.Format("2006-01-02")

	tablename := "account"
	sql := fmt.Sprintf("select ta.plat, count(ta.account_id) as accountnum, count(distinct ta.mac) as devicenum from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') = '%s'", startday)
	sql += fmt.Sprintf(" and not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	sql += fmt.Sprintf(" group by ta.plat;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getnewlyplatdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		platstr := utils.GetString("plat", v)
		results[platstr] = &base.PlatDnuData{
			Dnu:     utils.GetInt64("accountnum", v),
			RoleDnu: 0,
			MacDnu:  utils.GetInt64("devicenum", v),
		}
	}

	tablename = "roleinfo"
	sql = fmt.Sprintf("select ta.plat, count(ta.role_id) as rolenum from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') = '%s'", startday)
	sql += fmt.Sprintf(" and not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	sql += fmt.Sprintf(" group by ta.plat;")
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getnewlyplatdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}

	for _, v := range buckets {
		platstr := utils.GetString("plat", v)
		result, ok := results[platstr]
		if ok {
			result.RoleDnu = utils.GetInt64("rolenum", v)
		} else {
			results[platstr] = &base.PlatDnuData{
				Dnu:     0,
				RoleDnu: utils.GetInt64("rolenum", v),
				MacDnu:  0,
			}
		}

	}

	return nil
}

func getplatremain(engine *xorm.Engine, realtime time.Time, tm time.Time, results map[string]*base.PlatRemainData) error {
	faygo.Debugf("realtime:(%v) tm:(%v)", realtime, tm)

	tablename := reader.gettablename("accountlogin", realtime)
	sql := fmt.Sprintf("select a.plat, count(distinct b.account_id) as accountnum, count(distinct b.mac) as devicenum from account as a inner join %s as b", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(a.created_time, '%%Y-%%m-%%d') = '%s' and a.account_id = b.account_id", tm.Format("2006-01-02"))
	sql += fmt.Sprintf(" and not exists (select 1 from log_filter as tf where tf.account_id=a.account_id )")
	sql += fmt.Sprintf(" group by a.plat;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getplatremain >>>> DBEngine Query sql Error :", err)
		return err
	}

	for _, v := range buckets {
		platstr := utils.GetString("plat", v)
		results[platstr] = &base.PlatRemainData{
			Num:     utils.GetInt64("accountnum", v),
			RoleNum: 0,
			MacNum:  utils.GetInt64("devicenum", v),
		}
	}

	tablename = reader.gettablename("rolelogin", realtime)
	sql = fmt.Sprintf("select a.plat, count(a.role_id) as rolenum from roleinfo as a inner join %s as b", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(a.created_time, '%%Y-%%m-%%d') = '%s' and a.role_id = b.role_id", tm.Format("2006-01-02"))
	sql += fmt.Sprintf(" and not exists (select 1 from log_filter as tf where tf.account_id=a.account_id )")
	sql += fmt.Sprintf(" group by a.plat;")
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getplatremain >>>> DBEngine Query sql Error :", err)
		return err
	}

	for _, v := range buckets {
		platstr := utils.GetString("plat", v)
		result, ok := results[platstr]
		if ok {
			result.RoleNum = utils.GetInt64("rolenum", v)
		} else {
			results[platstr] = &base.PlatRemainData{
				Num:     0,
				RoleNum: utils.GetInt64("rolenum", v),
				MacNum:  0,
			}
		}
	}

	return nil
}

func getplatdau(engine *xorm.Engine, realtime time.Time, results map[string]*base.PlatDauData) error {
	//timestr := realtime.Format("2006-01-02")
	tablename := reader.gettablename("accountlogin", realtime)
	sql := fmt.Sprintf("select ta.plat, count(distinct ta.account_id) as accountnum, count(distinct ta.mac) as devicenum from %s as ta", tablename)
	sql += fmt.Sprintf(" where not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	sql += fmt.Sprintf(" group by ta.plat;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getplatdau >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		platstr := utils.GetString("plat", v)
		results[platstr] = &base.PlatDauData{
			Dau:     utils.GetInt64("accountnum", v),
			RoleDau: 0,
			MacDau:  utils.GetInt64("devicenum", v),
		}
	}

	tablename = reader.gettablename("rolelogin", realtime)
	sql = fmt.Sprintf("select ta.plat, count(ta.role_id) as rolenum from %s as ta", tablename)
	sql += fmt.Sprintf(" where not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	sql += fmt.Sprintf(" group by ta.plat;")
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getplatdau >>>> DBEngine Query sql Error : %v", err)
		return err
	}

	for _, v := range buckets {
		platstr := utils.GetString("plat", v)

		result, ok := results[platstr]
		if ok {
			result.RoleDau = utils.GetInt64("rolenum", v)
		} else {
			results[platstr] = &base.PlatDauData{
				Dau:     0,
				RoleDau: utils.GetInt64("rolenum", v),
				MacDau:  0,
			}
		}
	}
	return nil
}

func getrunoff(engine *xorm.Engine, tm time.Time, querytype string, daycounter int, series *[]*statistics.RunOffData) error {
	day := fmt.Sprintf("%s", tm.Format("2006-01-02"))
	offset := time.Hour * time.Duration(24*daycounter)

	atablename := "roleinfo"
	btablename := reader.gettablename("rolelogin", tm.Add(offset))
	tablename := reader.gettablename("rolelogout", tm)
	sql := fmt.Sprintf("select a.plat, c.%s as keystr, count(1) as data from %s as a left join %s as b on a.account_id=b.account_id and a.role_id=b.role_id",
		querytype, atablename, btablename)
	sql += fmt.Sprintf(" inner join %s as c", tablename)
	sql += fmt.Sprintf(" where not exists (select 1 from log_filter as tf where tf.account_id=a.account_id )")
	sql += fmt.Sprintf(" and DATE_FORMAT(a.created_time, '%%Y-%%m-%%d') = '%s' ", day)
	sql += fmt.Sprintf(" and b.account_id is null and b.role_id is null and a.account_id=c.account_id and a.role_id=c.role_id")
	sql += fmt.Sprintf(" group by a.plat, c.%s;", querytype)
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getrunoff >>>> DBEngine Query sql Error : %v", err)
		return err
	}

	for _, v := range buckets {
		plat := utils.GetInt("plat", v)
		key := utils.GetString("keystr", v)
		val := utils.GetInt64("data", v)

		*series = append(*series, &statistics.RunOffData{
			Plat:    plat,	
			Number:  key,
			RoleNum: val,
		})
	}

	return nil
}
