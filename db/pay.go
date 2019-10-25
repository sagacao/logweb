package db

import (
	"fmt"
	"logweb/ini"
	"logweb/models/cost"
	"logweb/models/pay"
	"sort"
	"time"

	"github.com/henrylee2cn/faygo"
)

// GetPayLevelData pay level
func GetPayLevelData(zeusid string, stm, etm time.Time, results *[]*pay.PayLevel) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	paylevel := make(map[int]*pay.PayLevel)
	getpayleveldata(engine, stm, etm, paylevel)

	var keys []int
	for k := range paylevel {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		*results = append(*results, paylevel[key])
	}

	return nil
}

// GetPayHabitData pay habit
func GetPayHabitData(zeusid string, habittype int, stm, etm time.Time, results *[]*pay.RechargePrice) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	payhabit := make(map[int]*pay.RechargePrice)
	getpayhabitdata(engine, habittype, stm, etm, payhabit)

	var keys []int
	for k := range payhabit {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		*results = append(*results, payhabit[key])
	}

	return nil
}

//GetPayHourData hour pay data
func GetPayHourData(zeusid string, stm, etm time.Time, results *[]*pay.PayHour) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	payhour := make(map[string]*pay.PayHour)
	getpayhourdata(engine, stm, etm, payhour)

	var keys []string
	for k := range payhour {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		*results = append(*results, payhour[key])
	}

	return nil
}

//GetPayRankData all change info
func GetPayRankData(zeusid string, istimerange int, stm, etm time.Time, results *[]*pay.PayRank) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	if istimerange == 1 {
		getpaytimerangerankdata(engine, stm, etm, results)
	} else {
		getpayrankdata(engine, results)
	}

	return nil
}

// GetPayRankDayData every day
func GetPayRankDayData(zeusid string, stm time.Time, results *[]*pay.PayRank) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	getpayrankdaydata(engine, stm, results)

	return nil
}

//GetPayDailyData daily pay data
func GetPayDailyData(zeusid string, stm, etm time.Time, results *[]*pay.PayDaily) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	getpaydailydata(engine, stm, etm, results)

	return nil
}

// GetPayChannelData is search pay channel info
func GetPayChannelData(zeusid string, stm, etm time.Time, results *[]*pay.PayChannel) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	getpaychanneldata(engine, stm, etm, results)

	return nil
}

// GetPayDaysAmountData is search days amount info
func GetPayDaysAmountData(zeusid string, stm, etm time.Time, results *[]*pay.DaysPayAmount) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	ltv := [9]int{3, 7, 14, 30, 60, 90, 120, 150, 180}
	//payamounts := make(map[string]*pay.DaysPayAmount)
	for stm.Before(etm) {
		for _, v := range ltv {

			levtm := stm.AddDate(0, 0, v)
			getpaydaysamountdata(engine, v, stm, stm, levtm, results)
		}

		stm = stm.Add(time.Hour * 24)
	}

	return nil
}

// GetPayDaysLossData is search days loss info
func GetPayDaysLossData(zeusid string, stm, etm time.Time, results *[]*pay.PayLoss) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("db engine '%s' not exist", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	//payamounts := make(map[string]*pay.DaysPayAmount)
	for stm.Before(etm) {

		stm = stm.Add(time.Hour * 24)
	}

	return nil
}

// Getrecharge is search rechange info
func Getrecharge(stm, etm time.Time, results *[]*cost.RechargeSummary) error {

	for k, v := range ini.ZeusType.Info {
		getGetrecharge(k, v, stm, etm, results)
	}

	return nil
}

// GetNewlyrecharge is search newly rechange info
func GetNewlyrecharge(stm, etm time.Time, results *[]*cost.RechargeSummary) error {

	for k, v := range ini.ZeusType.Info {
		getNewlyRecharge(k, v, stm, etm, results)
	}

	return nil
}

// GetHourRecharge is search hour rechange info
func GetHourRecharge(stm, etm time.Time, results *[]*cost.RechargeSummary) error {

	for k, v := range ini.ZeusType.Info {
		getHourRecharge(k, v, stm, etm, results)
	}

	return nil
}
