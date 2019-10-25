package handler

import (
	"fmt"
	"logweb/db"
	"logweb/models/base"
	"logweb/models/statistics"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

/*
RemainData
*/
var RemainData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("RemainData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	remaindata := make([]*statistics.RemainData, 0)
	statistics.InitRemainData(stm, etm, &remaindata)
	err := db.GetRemainData(zeusid, stm, etm, &remaindata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	faygo.Infof("RemainData ------------------------- ")
	for _, v := range remaindata {
		pag.Push(v)
		faygo.Infof("RemainData %v", v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
RemainPlatData
*/
var RemainPlatData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour * 7)
	faygo.Infof("RemainPlatData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	//remaindata := make([]*statistics.RemainData, 0)
	//statistics.InitRemainData(stm, etm, &remaindata)
	platremainretdata := make(map[string]*statistics.RemainData)
	err := db.GetPlatRemainData(zeusid, stm, etm, platremainretdata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range platremainretdata {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
RemainModelData
*/
var RemainModelData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("RemainModelData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	//payLevel := make([]*pay.PayLevel, 0)
	//err := db.GetPayLevelData(zeusid, stm, etm, &payLevel)
	//if err != nil {
	//	pag.Warnings = fmt.Sprintf("%v", err)
	//}
	return ctx.JSONMsg(200, 200, pag)
})

/*
RegisterData
*/
var RegisterData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("RegisterData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	//payLevel := make([]*pay.PayLevel, 0)
	//err := db.GetPayLevelData(zeusid, stm, etm, &payLevel)
	//if err != nil {
	//	pag.Warnings = fmt.Sprintf("%v", err)
	//}
	return ctx.JSONMsg(200, 200, pag)
})

/*
DauData
*/
var DauData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("DauData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	daydau := make([]*base.DayDauData, 0)
	err := db.GetDauData(zeusid, stm, etm, &daydau)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}
	for _, v := range daydau {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
DauPlatData
*/
var DauPlatData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("DauData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	platdau := make([]*base.PlatDauData, 0)
	err := db.GetDauPlatData(zeusid, stm, etm, &platdau)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range platdau {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

/*
RunoffData
*/
var RunoffData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("RunoffData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	querytype := ctx.QueryParam("runofftype")
	daycounter := ctx.QueryParam("para")
	faygo.Infof("RunoffData querytype:[%s] daycounter:[%v]", querytype, daycounter)
	runoffdata := make([]*statistics.RunOffData, 0)
	err := db.GetRunoffData(zeusid, stm, etm, querytype, utils.ParseInt(daycounter), &runoffdata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range runoffdata {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})
