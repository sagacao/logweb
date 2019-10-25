package db

import (
	"fmt"
	"logweb/models/base"
	"logweb/utils"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/henrylee2cn/faygo"
)

func getnewlydata(engine *xorm.Engine, stm, etm time.Time, results map[string]*base.DayDnuData) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "account"
	sql := fmt.Sprintf("select DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') as tmstr, count(distinct ta.account_id) as accountnum, count(distinct ta.mac) as devicenum from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') <= '%s'", startday, endday)
	sql += fmt.Sprintf(" and not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	sql += fmt.Sprintf(" group by tmstr order by tmstr;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getnewlydata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		timestr := utils.GetString("tmstr", v)
		results[timestr] = &base.DayDnuData{
			Dnu:     utils.GetInt64("accountnum", v),
			RoleDnu: 0,
			MacDnu:  utils.GetInt64("devicenum", v),
		}
	}

	tablename = "roleinfo"
	sql = fmt.Sprintf("select DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') as tmstr, count(ta.role_id) as rolenum from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') <= '%s'", startday, endday)
	sql += fmt.Sprintf(" and not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id )")
	sql += fmt.Sprintf(" group by tmstr order by tmstr;")
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getnewlydata >>>> DBEngine Query sql Error : %v", err)
		return err
	}

	for _, v := range buckets {
		timestr := utils.GetString("tmstr", v)
		result, ok := results[timestr]
		if ok {
			result.RoleDnu = utils.GetInt64("rolenum", v)
		} else {
			results[timestr] = &base.DayDnuData{
				Dnu:     0,
				RoleDnu: utils.GetInt64("rolenum", v),
				MacDnu:  0,
			}
		}
	}

	return nil
}

func getdailydau(engine *xorm.Engine, realtime time.Time, results map[string]*base.DayDauData) error {
	timestr := realtime.Format("2006-01-02")
	tablename := reader.gettablename("accountlogin", realtime)
	sql := fmt.Sprintf("select count(distinct ta.account_id) as accountnum, count(distinct ta.mac) as devicenum from %s as ta", tablename)
	sql += fmt.Sprintf(" where not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id );")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailydau >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		results[timestr] = &base.DayDauData{
			Dau:     utils.GetInt64("accountnum", v),
			RoleDau: 0,
			MacDau:  utils.GetInt64("devicenum", v),
		}
	}

	tablename = reader.gettablename("rolelogin", realtime)
	sql = fmt.Sprintf("select count(ta.role_id) as rolenum from %s as ta", tablename)
	sql += fmt.Sprintf(" where not exists (select 1 from log_filter as tf where tf.account_id=ta.account_id );")
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailydau >>>> DBEngine Query sql Error : %v", err)
		return err
	}

	for _, v := range buckets {
		result, ok := results[timestr]
		if ok {
			result.RoleDau = utils.GetInt64("rolenum", v)
		} else {
			results[timestr] = &base.DayDauData{
				Dau:     0,
				RoleDau: utils.GetInt64("rolenum", v),
				MacDau:  0,
			}
		}
	}
	return nil
}

func getdailyremain(engine *xorm.Engine, realtime time.Time, tm time.Time, results map[string]*base.DayRemainData) error {
	faygo.Debugf("realtime:(%v) tm:(%v)", realtime, tm)
	timestr := tm.Format("2006-01-02")
	tablename := reader.gettablename("accountlogin", realtime)
	sql := fmt.Sprintf("select count(distinct b.account_id) as accountnum, count(distinct b.mac) as devicenum from account as a inner join %s as b", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(a.created_time, '%%Y-%%m-%%d') = '%s' and a.account_id = b.account_id", timestr)
	sql += fmt.Sprintf(" and not exists (select 1 from log_filter as tf where tf.account_id=a.account_id );")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyremain >>>> DBEngine Query sql Error :", err)
		return err
	}

	for _, v := range buckets {
		results[timestr] = &base.DayRemainData{
			Num:     utils.GetInt64("accountnum", v),
			RoleNum: 0,
			MacNum:  utils.GetInt64("devicenum", v),
		}
	}

	tablename = reader.gettablename("rolelogin", realtime)
	sql = fmt.Sprintf("select count(1) as rolenum from roleinfo as a inner join %s as b", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(a.created_time, '%%Y-%%m-%%d') = '%s' and a.role_id = b.role_id", timestr)
	sql += fmt.Sprintf(" and not exists (select 1 from log_filter as tf where tf.account_id=a.account_id );")
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyremain >>>> DBEngine Query sql Error :", err)
		return err
	}

	for _, v := range buckets {
		result, ok := results[timestr]
		if ok {
			result.RoleNum = utils.GetInt64("rolenum", v)
		} else {
			results[timestr] = &base.DayRemainData{
				Num:     0,
				RoleNum: utils.GetInt64("rolenum", v),
				MacNum:  0,
			}
		}
	}

	return nil
}

func getdailyrecharge(engine *xorm.Engine, realtime time.Time, results map[string]*base.DayRechargeData) error {
	timestr := realtime.Format("2006-01-02")

	tablename := reader.gettablename("chargeinfo", realtime)
	sql := fmt.Sprintf("select sum(ta.cash) as totalcash, count(distinct ta.role_id) as rechargeuser from %s as ta", tablename)
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		results[timestr] = &base.DayRechargeData{
			TotalIncome:   utils.GetFloat32("totalcash", v),
			RechargeUsers: utils.GetInt64("rechargeuser", v),
		}
	}
	return nil
}

func getdailynewrecharge(engine *xorm.Engine, realtime time.Time, results map[string]*base.DayNewRechargeData) error {
	timestr := realtime.Format("2006-01-02")

	tablename := reader.gettablename("chargeinfo", realtime)
	sql := fmt.Sprintf("select sum(ta.cash) as totalcash, count(distinct ta.role_id) as rechargeuser from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.reg_time, '%%Y-%%m-%%d') = '%s'", realtime.Format("2006-01-02"))
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailynewrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		results[timestr] = &base.DayNewRechargeData{
			NewRechargeIncome: utils.GetFloat32("totalcash", v),
			NewRechargeUsers:  utils.GetInt64("rechargeuser", v),
		}
	}
	return nil
}
