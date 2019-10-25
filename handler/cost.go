package handler

import (
	"fmt"
	"logweb/db"
	"logweb/models/base"
	"logweb/models/cost"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

/*
MoneyChangePage
*/
var MoneyChangePage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("MoneyChangePage zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	costtype := ctx.QueryParam("type")
	moneytype := ctx.QueryParam("moneyType")
	faygo.Infof("costtype:[%s] moneytype:[%s]", costtype, moneytype)

	moneyChange := make([]*cost.MoneyChange, 0)
	err := db.GetMoneyChangePage(zeusid, utils.StringToInt(costtype), utils.StringToInt(moneytype), stm, etm, &moneyChange)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range moneyChange {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
MoneyChangeList
*/
var MoneyChangeList = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("MoneyChangeList zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	moneyChangeList := make([]*cost.MoneyChangeSummary, 0)
	cost.InitMoneyChangeSummary(stm, etm, &moneyChangeList)

	err := db.GetMoneyChangeList(zeusid, stm, etm, &moneyChangeList)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range moneyChangeList {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

/*
MallCostPage
*/
var MallCostPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("MallCostPage zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	//payLevel := make([]*pay.PayLevel, 0)
	//err := db.GetPayLevelData(zeusid, stm, etm, &payLevel)
	//if err != nil {
	//	pag.Warnings = fmt.Sprintf("%v", err)
	//}
	return ctx.JSONMsg(200, 200, pag)
})

/*
MoneyChangeList
*/
var MallCostList = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("MallCostList zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	//payLevel := make([]*pay.PayLevel, 0)
	//err := db.GetPayLevelData(zeusid, stm, etm, &payLevel)
	//if err != nil {
	//	pag.Warnings = fmt.Sprintf("%v", err)
	//}
	return ctx.JSONMsg(200, 200, pag)
})
