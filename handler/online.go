package handler

import (
	"fmt"
	"logweb/db"
	"logweb/models/base"
	"time"

	"github.com/henrylee2cn/faygo"
)

/*
Online
*/
var Online = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))

	faygo.Infof("Online zeusid:[%s] time:[%v]", zeusid, stm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	onlinedata := make(map[int]*base.OnlineData)
	err := db.GetOnlinedata(zeusid, stm, onlinedata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	faygo.Infof("Online ------------------------- ")
	for _, v := range onlinedata {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
CurOnline
*/
var CurOnline = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	faygo.Infof("CurOnline time:[%v]", time.Now())
	
	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	curOnlineData := make(map[string]*base.CurOnlineData)
	err := db.GetCurOnline(curOnlineData)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	faygo.Infof("CurOnline ------------------------- ")
	for _, v := range curOnlineData {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
AvgOnline
*/
var AvgOnline = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	prof := ctx.QueryParam("prof")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("AvgOnline zeusid:[%s] timerange:[%v || %v]  prof:[%v]", zeusid, stm, etm, prof)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	// avgOnlineData := make(map[string]*base.AvgOnlineData)
	avgOnlineData := make([]*base.AvgOnlineData, 0)
	err := db.GetAvgOnlineData(zeusid, prof, stm, etm, &avgOnlineData)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	faygo.Infof("AvgOnline ------------------------- ")
	for _, v := range avgOnlineData {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
LeagueLine
*/
var LeagueLine = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))

	faygo.Infof("LeagueLine zeusid:[%s] time:[%v]", zeusid, stm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	leagueLineData := make(map[int]*base.LeagueLineData)
	err := db.GetLeagueLineData(zeusid, stm, leagueLineData)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	faygo.Infof("LeagueLine ------------------------- ")
	for _, v := range leagueLineData {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})
