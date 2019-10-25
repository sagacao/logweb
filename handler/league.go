package handler

import (
	"fmt"
	"logweb/db"
	"logweb/models/base"
	"logweb/models/league"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

/*
LeagueData
*/
var LeagueData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("LeagueData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	prof := ctx.QueryParam("prof")
	faygo.Infof("LeagueData prof:[%v]", prof)

	leaguedata := make([]*league.LeagueData, 0)
	err := db.GetLeagueData(zeusid, stm, etm, &leaguedata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range leaguedata {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

var LeagueInfo = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))

	faygo.Infof("LeagueInfo zeusid:[%s] time:[%v]", zeusid, stm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	number := ctx.QueryParam("number")
	faygo.Infof("LeagueInfo number:[%v]", number)

	leagueinfo := make([]*league.LeagueInfo, 0)
	err := db.GetLeagueInfo(zeusid, stm, utils.StringToInt(number), &leagueinfo)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	leaguemap := make(map[string][]int64)
	for _, v := range leagueinfo {
		leaguemap["level"] = append(leaguemap["level"], v.Level)
		leaguemap["assets"] = append(leaguemap["assets"], v.Assets)
		leaguemap["credits"] = append(leaguemap["credits"], v.Credits)
		leaguemap["health"] = append(leaguemap["health"], v.Health)
		leaguemap["member"] = append(leaguemap["member"], v.Member)
		leaguemap["student"] = append(leaguemap["student"], v.Student)
	}
	pag.Push(leaguemap)

	return ctx.JSONMsg(200, 200, pag)
})
