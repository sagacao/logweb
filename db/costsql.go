package db

import (
	"fmt"
	"logweb/models/cost"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"

	"github.com/go-xorm/xorm"
)

func getmoneychangepage(engine *xorm.Engine, costtype int, moneytype int, stm, etm time.Time, series *[]*cost.MoneyChange) error {

	sql := fmt.Sprintf("select reason, sum(val) as num from (")
	counter := 0
	var tablename string
	for stm.Before(etm) {
		if counter > 0 {
			sql += fmt.Sprintf(" union all ")
		}
		if costtype == 1 {
			tablename = reader.gettablename("moneyproduct", stm)
		}
		if costtype == -1 {
			tablename = reader.gettablename("moneyconsume", stm)
		}
		sql += fmt.Sprintf("(select reason, sum(value)  as val from %s where money_type=%d group by reason)", tablename, moneytype)

		counter++
		stm = stm.Add(time.Hour * 24)
	}
	sql += fmt.Sprintf(") as tab_all group by reason")
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getnewlyplatdata >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		*series = append(*series, &cost.MoneyChange{
			Reason: utils.GetString("reason", v),
			Num:    utils.GetInt64("num", v),
		})
	}

	return nil
}

func getdailymoney(engine *xorm.Engine, realtime time.Time, results map[string]*cost.Moneysummary) error {
	timestr := realtime.Format("2006-01-02")

	tablename := reader.gettablename("moneyproduct", realtime)
	sql := fmt.Sprintf("select sum(value) as val from %s where money_type=1", tablename)
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		results[timestr] = &cost.Moneysummary{
			MoneyGain: utils.GetInt64("val", v),
			MoneyCost: 0,
		}
	}

	tablename = reader.gettablename("moneyconsume", realtime)
	sql = fmt.Sprintf("select sum(value) as val from %s where money_type=1", tablename)
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		result, ok := results[timestr]
		if ok {
			result.MoneyCost = utils.GetInt64("val", v)
		} else {
			results[timestr] = &cost.Moneysummary{
				MoneyGain: 0,
				MoneyCost: utils.GetInt64("val", v),
			}
		}
	}

	return nil
}

func getdailydimmi(engine *xorm.Engine, realtime time.Time, results map[string]*cost.Dimmisummary) error {
	timestr := realtime.Format("2006-01-02")

	tablename := reader.gettablename("moneyproduct", realtime)
	sql := fmt.Sprintf("select sum(value) as val from %s where money_type=4", tablename)
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		results[timestr] = &cost.Dimmisummary{
			DimmiGain: utils.GetInt64("val", v),
			DimmiCost: 0,
		}
	}

	tablename = reader.gettablename("moneyconsume", realtime)
	sql = fmt.Sprintf("select sum(value) as val from %s where money_type=4", tablename)
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		result, ok := results[timestr]
		if ok {
			result.DimmiCost = utils.GetInt64("val", v)
		} else {
			results[timestr] = &cost.Dimmisummary{
				DimmiGain: 0,
				DimmiCost: utils.GetInt64("val", v),
			}
		}
	}

	return nil
}

func getdailygold(engine *xorm.Engine, realtime time.Time, results map[string]*cost.Goldsummary) error {
	timestr := realtime.Format("2006-01-02")

	tablename := reader.gettablename("moneyproduct", realtime)
	sql := fmt.Sprintf("select sum(value) as val from %s where money_type=3", tablename)
	printsql(sql)

	buckets, err := engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		results[timestr] = &cost.Goldsummary{
			GoldGain: utils.GetInt64("val", v),
			GoldCost: 0,
		}
	}

	tablename = reader.gettablename("moneyconsume", realtime)
	sql = fmt.Sprintf("select sum(value) as val from %s where money_type=3", tablename)
	printsql(sql)

	buckets, err = engine.Query(sql)
	if err != nil && !utils.CatchError("Error 1146", err) {
		faygo.Errorf("getdailyrecharge >>>> DBEngine Query sql Error : %v", err)
		return err
	}
	printsql(buckets)

	for _, v := range buckets {
		result, ok := results[timestr]
		if ok {
			result.GoldCost = utils.GetInt64("val", v)
		} else {
			results[timestr] = &cost.Goldsummary{
				GoldGain: 0,
				GoldCost: utils.GetInt64("val", v),
			}
		}
	}

	return nil
}
