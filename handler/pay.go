package handler

import (
	"fmt"
	"logweb/db"
	"logweb/models/base"
	"logweb/models/cost"
	"logweb/models/pay"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

/*
PayLevelData
*/
var PayLevelData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("PayLevelData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	payLevel := make([]*pay.PayLevel, 0)
	err := db.GetPayLevelData(zeusid, stm, etm, &payLevel)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}
	for _, v := range payLevel {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
PayHabitData
*/
var PayHabitData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("PayHabitData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	habitType := ctx.QueryParam("habitType")
	payhabit := make([]*pay.RechargePrice, 0)
	db.GetPayHabitData(zeusid, utils.StringToInt(habitType), stm, etm, &payhabit)

	for _, v := range payhabit {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
PayHourData
*/
var PayHourData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("PayHourData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	payhour := make([]*pay.PayHour, 0)
	db.GetPayHourData(zeusid, stm, etm, &payhour)
	for _, v := range payhour {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

/*
PayAmount //用户贡献
*/
var PayAmountData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("PayAmountData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	paydata := make([]*pay.DaysPayAmount, 0)
	pay.InitDaysPayAmount(stm, etm, &paydata)
	db.GetPayDaysAmountData(zeusid, stm, etm, &paydata)
	for _, v := range paydata {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

/*
PayRankData
*/
var PayRankData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("PayRankData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	rankType := ctx.QueryParam("rankType")
	payrank := make([]*pay.PayRank, 0)
	db.GetPayRankData(zeusid, utils.StringToInt(rankType), stm, etm, &payrank)
	for _, v := range payrank {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

/*
VipData
*/
var VipData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("VipData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	//rankType := ctx.QueryParam("rankType")
	payrank := make([]*pay.PayRank, 0)
	//db.GetPayRankData("*", utils.StringToInt(rankType), stm, etm, &payrank)

	return ctx.JSONMsg(200, 200, payrank)
})

/*
PayDailyData
*/
var PayDailyData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("PayDailyData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	paydata := make([]*pay.PayDaily, 0)
	db.GetPayDailyData(zeusid, stm, etm, &paydata)
	for _, v := range paydata {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

/*
PayChannelData
*/
var PayChannelData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("PayChannelData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	paydata := make([]*pay.PayChannel, 0)
	db.GetPayChannelData(zeusid, stm, etm, &paydata)
	for _, v := range paydata {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

/*
PayLossData
*/
var PayLossData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("PayLossData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	paydata := make([]*pay.PayLoss, 0)
	db.GetPayDaysLossData(zeusid, stm, etm, &paydata)
	for _, v := range paydata {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

//Recharge recharge sum
var Recharge = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("timerange:[%v || %v]", stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	rechargedata := make([]*cost.RechargeSummary, 0)
	err := db.Getrecharge(stm, etm, &rechargedata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range rechargedata {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)

})

//NewlyRecharge newly recharge sum
var NewlyRecharge = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("timerange:[%v || %v]", stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	rechargedata := make([]*cost.RechargeSummary, 0)
	err := db.GetNewlyrecharge(stm, etm, &rechargedata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range rechargedata {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)

})

// HourRecharge hour sum data
var HourRecharge = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("timerange:[%v || %v]", stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	rechargedata := make([]*cost.RechargeSummary, 0)
	err := db.GetHourRecharge(stm, etm, &rechargedata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range rechargedata {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)

})

/*
PayRankDayData every day pay rank
*/
var PayRankDayData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.AddDate(0, 0, 1)

	faygo.Infof("PayRankDayData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	payrank := make([]*pay.PayRank, 0)
	db.GetPayRankDayData(zeusid, stm, &payrank)
	for _, v := range payrank {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})
