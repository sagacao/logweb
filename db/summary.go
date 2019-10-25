package db

import (
	"fmt"
	"logweb/models/base"
	"logweb/models/summary"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

func GetSummaryData(zeusid string, stm, etm time.Time, results *[]*summary.GameSummary) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}
	dnu := make(map[string]*base.DayDnuData)
	getnewlydata(engine, stm, etm, dnu)
	
	dayrecharge := make(map[string]*base.DayRechargeData)
	dau := make(map[string]*base.DayDauData)
	daynewrecharge := make(map[string]*base.DayNewRechargeData)
	for stm.Before(etm) {
		getdailyrecharge(engine, stm, dayrecharge)
		getdailydau(engine, stm, dau)
		getdailynewrecharge(engine, stm, daynewrecharge)
		stm = stm.Add(time.Hour * 24)
	}

	for _, v := range *results {
		key := v.DimDay
		dnudata, ok := dnu[key]
		if ok {
			v.Dnu = dnudata.Dnu
			v.MacDnu = dnudata.MacDnu
		}

		daudata, ok := dau[key]
		if ok {
			v.Dau = daudata.Dau
			v.MacDau = daudata.MacDau
		}
		
		rechargedata, ok := dayrecharge[key]
		if ok {
			v.TotalIncome = rechargedata.TotalIncome
			v.RechargeUsers = rechargedata.RechargeUsers
		}

		newrechargedata, ok := daynewrecharge[key]
		if ok {
			v.NewRechargeIncome = newrechargedata.NewRechargeIncome
			v.NewRechargeUsers = newrechargedata.NewRechargeUsers
		}

		v.NewPayRate = utils.FormatPercent(float32(v.NewRechargeUsers), float32(v.Dnu))
		v.PayRate = utils.FormatPercent(float32(v.RechargeUsers), float32(v.Dau))
		v.Arppu = utils.FormatPercent(float32(v.TotalIncome), float32(v.RechargeUsers))
		v.WithoutNewDau = int32(v.Dau - v.Dnu)
	}

	return nil
}
