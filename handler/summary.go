package handler

import (
	"fmt"
	"logweb/db"
	"logweb/models/base"
	"logweb/models/summary"
	"time"

	"github.com/henrylee2cn/faygo"
)

/*
SummaryData
*/
var SummaryData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("SummaryData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	gameSummary := make([]*summary.GameSummary, 0)
	summary.InitGameSummary(stm, etm, &gameSummary)

	err := db.GetSummaryData(zeusid, stm, etm, &gameSummary)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
ModelData
*/
var ModelData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	return ctx.Render(200, faygo.JoinStatic("index.html"), faygo.Map{
		"TITLE":   "faygo",
		"VERSION": faygo.VERSION,
		"CONTENT": "Welcome To Faygo",
		"AUTHOR":  "HenryLee",
	})
})
