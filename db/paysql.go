package db

import (
	"fmt"
	"logweb/models/cost"
	"logweb/models/pay"
	"logweb/utils"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/henrylee2cn/faygo"
)

func getpayleveldata(engine *xorm.Engine, stm, etm time.Time, results map[int]*pay.PayLevel) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select ta.level, count(distinct ta.role_id) as payercount, count(1) as paycount, sum(ta.cash) as amount from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" group by ta.level;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpayleveldata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		level := utils.GetInt("level", v)

		results[level] = &pay.PayLevel{
			Level:      level,
			Amount:     utils.GetFloat32("amount", v),
			PayerCount: utils.GetInt64("payercount", v),
			PayCount:   utils.GetInt64("paycount", v),
		}
	}

	return nil
}

func getpayhabitdata(engine *xorm.Engine, habittype int, stm, etm time.Time, results map[int]*pay.RechargePrice) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select ta.recharge_id, ta.cash, count(distinct ta.role_id) as payercount, count(1) as paycount from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	if habittype != 0 {
		sql += fmt.Sprintf(" and recharge_type = %d", habittype)
	}
	sql += fmt.Sprintf(" group by ta.recharge_id;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpayhabitdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		price := utils.GetInt("cash", v)
		charge := utils.GetInt("recharge_id", v)

		results[charge] = &pay.RechargePrice{
			ChargeID:   charge,
			Price:      price,
			PayerCount: utils.GetInt64("payercount", v),
			PayCount:   utils.GetInt64("paycount", v),
		}
	}

	return nil
}

func getpayhourdata(engine *xorm.Engine, stm, etm time.Time, results map[string]*pay.PayHour) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select DATE_FORMAT(ta.log_time, '%%H') as hour, count(distinct ta.role_id) as payercount, sum(ta.cash) as amount from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" group by hour;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpayhabitdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		hour := utils.GetString("hour", v)

		results[hour] = &pay.PayHour{
			Hour:       hour,
			PayerCount: utils.GetInt64("payercount", v),
			Amount:     utils.GetFloat32("amount", v),
		}
	}

	return nil
}

func getpayrankdata(engine *xorm.Engine, results *[]*pay.PayRank) error {

	tablename := "chargeinfo"
	//sql := fmt.Sprintf("select max(ta.total_cash) as amount, ta.role_id, ta.name as name, max(ta.level) as level, min(ta.reg_time) as reg_time, ta.plat as plat, ta.server_id as server_id, ta.os as os, max(ta.log_time) as log_time from %s as ta", tablename)
	//sql += fmt.Sprintf(" group by ta.role_id,name order by amount, log_time desc limit 300;")
	sql := fmt.Sprintf("select max(ta.total_cash) as amount, ta.role_id, ta.name, max(ta.level) as level, min(ta.reg_time) as reg_time, ta.plat, ta.server_id, ta.os, ta.log_time")
	sql += fmt.Sprintf(" from (select max(log_time) as timestamp, role_id from %s group by role_id order by log_time desc) as tb", tablename)
	sql += fmt.Sprintf(" left join %s as ta on tb.role_id=ta.role_id and tb.timestamp=ta.log_time", tablename)
	sql += fmt.Sprintf(" group by ta.role_id order by amount desc limit 300;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpayrankdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	rank := 1
	for _, v := range buckets {
		*results = append(*results, &pay.PayRank{
			Rank:          rank,
			PlayerId:      utils.GetInt("role_id", v),
			Name:          utils.GetString("name", v),
			Level:         utils.GetInt("level", v),
			Amount:        utils.GetFloat32("amount", v),
			Gold:          0,
			Zeusid:        utils.GetInt("server_id", v),
			Channelid:     utils.GetString("plat", v),
			CreateTime:    utils.GetString("reg_time", v),
			LastLoginTime: utils.GetString("log_time", v),
		})

		rank++
	}

	return nil
}

func getpaytimerangerankdata(engine *xorm.Engine, stm, etm time.Time, results *[]*pay.PayRank) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "chargeinfo"
	//sql := fmt.Sprintf("select sum(ta.cash) as amount, ta.role_id, ta.name as name, max(ta.level) as level, min(ta.reg_time) as reg_time, ta.plat as plat, ta.server_id as server_id, ta.os as os, max(ta.log_time) as log_time from %s as ta", tablename)
	//sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	//sql += fmt.Sprintf(" group by ta.role_id,name order by amount desc limit 300;")
	sql := fmt.Sprintf("select max(ta.total_cash) as amount, ta.role_id, ta.name, max(ta.level) as level, min(ta.reg_time) as reg_time, ta.plat, ta.server_id, ta.os, ta.log_time")
	sql += fmt.Sprintf(" from (select max(log_time) as timestamp, role_id from %s ", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" group by role_id order by log_time desc) as tb")
	sql += fmt.Sprintf(" left join %s as ta on tb.role_id=ta.role_id and tb.timestamp=ta.log_time", tablename)
	sql += fmt.Sprintf(" group by ta.role_id order by amount desc limit 300;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpayrankdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	rank := 1
	for _, v := range buckets {
		*results = append(*results, &pay.PayRank{
			Rank:          rank,
			PlayerId:      utils.GetInt("role_id", v),
			Name:          utils.GetString("name", v),
			Level:         utils.GetInt("level", v),
			Amount:        utils.GetFloat32("amount", v),
			Gold:          0,
			Zeusid:        utils.GetInt("server_id", v),
			Channelid:     utils.GetString("plat", v),
			CreateTime:    utils.GetString("reg_time", v),
			LastLoginTime: utils.GetString("log_time", v),
		})

		rank++
	}

	return nil
}

func getpayrankdaydata(engine *xorm.Engine, stm time.Time, results *[]*pay.PayRank) error {
	startday := stm.Format("2006-01-02")

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select sum(ta.cash) as amount, ta.role_id, ta.name as name, max(ta.level) as level, min(ta.reg_time) as reg_time, ta.plat as plat, ta.server_id as server_id, ta.os as os, max(ta.log_time) as log_time from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') = '%s' ", startday)
	sql += fmt.Sprintf(" group by ta.role_id,name order by amount desc limit 300;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpayrankdaydata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	rank := 1
	for _, v := range buckets {
		*results = append(*results, &pay.PayRank{
			Rank:          rank,
			PlayerId:      utils.GetInt("role_id", v),
			Name:          utils.GetString("name", v),
			Level:         utils.GetInt("level", v),
			Amount:        utils.GetFloat32("amount", v),
			Gold:          0,
			Zeusid:        utils.GetInt("server_id", v),
			Channelid:     utils.GetString("plat", v),
			CreateTime:    utils.GetString("reg_time", v),
			LastLoginTime: utils.GetString("log_time", v),
		})

		rank++
	}

	return nil
}

func getpaydailydata(engine *xorm.Engine, stm, etm time.Time, results *[]*pay.PayDaily) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') as timestr, sum(ta.cash) as amount, count(distinct ta.role_id) as payercount from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" group by timestr order by timestr;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpaydailydata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		*results = append(*results, &pay.PayDaily{
			DimDay:     utils.GetString("timestr", v),
			PayerCount: utils.GetInt64("payercount", v),
			Amount:     utils.GetFloat32("amount", v),
		})
	}

	return nil
}

func getpaychanneldata(engine *xorm.Engine, stm, etm time.Time, results *[]*pay.PayChannel) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select ta.plat, sum(ta.cash) as amount, count(distinct ta.role_id) as payercount from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" group by ta.plat order by ta.plat;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpaychanneldata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		*results = append(*results, &pay.PayChannel{
			Channelid:  utils.GetInt("plat", v),
			PayerCount: utils.GetInt64("payercount", v),
			Amount:     utils.GetFloat32("amount", v),
		})
	}

	return nil
}

func getpaydaysamountdata(engine *xorm.Engine, ltvday int, realtm, stm, etm time.Time, results *[]*pay.DaysPayAmount) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	totalamount := float32(0)
	totalplayer := float32(1)

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select count(distinct ta.role_id) as newpayercount, sum(ta.cash) as amount from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" and DATE_FORMAT(ta.reg_time, '%%Y-%%m-%%d') = '%s';", startday)
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpaydaysamountdata >>>>1 DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)
	for _, v := range buckets {
		totalamount = utils.GetFloat32("amount", v)
		break
	}

	tablename = "roleinfo"
	sql = fmt.Sprintf("select count(distinct ta.role_id) as playercount from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.created_time, '%%Y-%%m-%%d') = '%s';", startday)
	playerbuckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getpaydaysamountdata >>>>2 DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(playerbuckets)
	for _, v := range playerbuckets {
		totalplayer = utils.GetFloat32("playercount", v)
		break
	}

	faygo.Debugf("getpaydaysamountdata amount=%04f, player=%04f", totalamount, totalplayer)
	dnp := float32(0)
	if totalplayer != 0 {
		dnp = float32(totalamount / totalplayer)
	}
	for _, rv := range *results {
		if rv.DimDay == startday {
			if ltvday == 3 {
				rv.Dnp3 = dnp
				rv.PayCount3 = int64(totalamount)
				rv.PlayerCount3 = int64(totalplayer)
			} else if ltvday == 7 {
				rv.Dnp7 = dnp
				rv.PayCount7 = int64(totalamount)
				rv.PlayerCount7 = int64(totalplayer)
			} else if ltvday == 14 {
				rv.Dnp14 = dnp
				rv.PayCount14 = int64(totalamount)
				rv.PlayerCount14 = int64(totalplayer)
			} else if ltvday == 30 {
				rv.Dnp30 = dnp
				rv.PayCount30 = int64(totalamount)
				rv.PlayerCount30 = int64(totalplayer)
			} else if ltvday == 60 {
				rv.Dnp60 = dnp
				rv.PayCount60 = int64(totalamount)
				rv.PlayerCount60 = int64(totalplayer)
			} else if ltvday == 90 {
				rv.Dnp90 = dnp
				rv.PayCount90 = int64(totalamount)
				rv.PlayerCount90 = int64(totalplayer)
			} else if ltvday == 120 {
				rv.Dnp120 = dnp
				rv.PayCount120 = int64(totalamount)
				rv.PlayerCount120 = int64(totalplayer)
			} else if ltvday == 150 {
				rv.Dnp150 = dnp
				rv.PayCount150 = int64(totalamount)
				rv.PlayerCount150 = int64(totalplayer)
			} else if ltvday == 180 {
				rv.Dnp180 = dnp
				rv.PayCount180 = int64(totalamount)
				rv.PlayerCount180 = int64(totalplayer)
			}
			break
		}
	}
	return nil
}

func getGetrecharge(zeusid string, zeusName string, stm, etm time.Time, results *[]*cost.RechargeSummary) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') as timestr,sum(ta.cash) as num from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >='%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" group by timestr;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getGetrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	var cashTotal float32
	for _, v := range buckets {
		cash := utils.GetFloat32("num", v)
		cashTotal = cash + cashTotal
		*results = append(*results, &cost.RechargeSummary{
			ServerName: zeusName,
			Time:       utils.GetString("timestr", v),
			Cash:       cash,
		})
	}

	*results = append(*results, &cost.RechargeSummary{
		ServerName: zeusName,
		Time:       "total",
		Cash:       cashTotal,
	})

	return nil
}

func getNewlyRecharge(zeusid string, zeusName string, stm, etm time.Time, results *[]*cost.RechargeSummary) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select DATE_FORMAT(ta.reg_time, '%%Y-%%m-%%d') as timestr, sum(ta.cash) as num from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.reg_time, '%%Y-%%m-%%d') >='%s' and DATE_FORMAT(ta.reg_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" and DATE_FORMAT(ta.reg_time, '%%Y-%%m-%%d') = DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d')")
	sql += fmt.Sprintf(" group by timestr;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getNewlyrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	var cashTotal float32
	for _, v := range buckets {
		cash := utils.GetFloat32("num", v)
		cashTotal = cash + cashTotal
		*results = append(*results, &cost.RechargeSummary{
			ServerName: zeusName,
			Time:       utils.GetString("timestr", v),
			Cash:       cash,
		})
	}

	*results = append(*results, &cost.RechargeSummary{
		ServerName: zeusName,
		Time:       "total",
		Cash:       cashTotal,
	})

	return nil
}

func getHourRecharge(zeusid string, zeusName string, stm, etm time.Time, results *[]*cost.RechargeSummary) error {
	startday := stm.Format("2006-01-02")
	endday := etm.Format("2006-01-02")

	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	tablename := "chargeinfo"
	sql := fmt.Sprintf("select DATE_FORMAT(ta.log_time, '%%H') as timestr, sum(ta.cash) as num from %s as ta", tablename)
	sql += fmt.Sprintf(" where DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') >= '%s' and DATE_FORMAT(ta.log_time, '%%Y-%%m-%%d') < '%s'", startday, endday)
	sql += fmt.Sprintf(" group by timestr;")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getHourRecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	var cashTotal float32
	for _, v := range buckets {
		cash := utils.GetFloat32("num", v)
		cashTotal = cash + cashTotal
		*results = append(*results, &cost.RechargeSummary{
			ServerName: zeusName,
			Time:       utils.GetString("timestr", v),
			Cash:       cash,
		})
	}

	*results = append(*results, &cost.RechargeSummary{
		ServerName: zeusName,
		Time:       "total",
		Cash:       cashTotal,
	})

	return nil
}
