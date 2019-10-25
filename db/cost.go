package db

import (
	"fmt"
	"logweb/models/cost"
	"time"

	"github.com/henrylee2cn/faygo"
)

func GetMoneyChangePage(zeusid string, costtype int, moneytype int, stm, etm time.Time, results *[]*cost.MoneyChange) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	getmoneychangepage(engine, costtype, moneytype, stm, etm, results)

	return nil
}

func GetMoneyChangeList(zeusid string, stm, etm time.Time, results *[]*cost.MoneyChangeSummary) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	moneysummary := make(map[string]*cost.Moneysummary, 0)
	dimmisummary := make(map[string]*cost.Dimmisummary, 0)
	goldsummary := make(map[string]*cost.Goldsummary, 0)
	for stm.Before(etm) {
		getdailymoney(engine, stm, moneysummary)
		getdailydimmi(engine, stm, dimmisummary)
		getdailygold(engine, stm, goldsummary)
		stm = stm.Add(time.Hour * 24)
	}

	for _, v := range *results {
		key := v.DimDay
		moneydata, ok := moneysummary[key]
		if ok {
			v.MoneyGain = moneydata.MoneyGain
			v.MoneyCost = moneydata.MoneyCost
			v.MoneyRemain = int64(v.MoneyGain - v.MoneyCost)
		}

		dimmidata, ok := dimmisummary[key]
		if ok {
			v.DimmiGain = dimmidata.DimmiGain
			v.DimmiCost = dimmidata.DimmiCost
			v.DimmiRemain = int64(v.DimmiGain - v.DimmiCost)
		}

		golddata, ok := goldsummary[key]
		if ok {
			v.GoldGain = golddata.GoldGain
			v.GoldCost = golddata.GoldCost
			v.GoldRemain = int64(v.GoldGain - v.GoldCost)
		}

	}

	return nil
}
