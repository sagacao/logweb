package handler

import (
	"fmt"
	"logweb/db"
	"logweb/models/base"
	"logweb/models/level"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

/*
LevelDisData
*/
var LevelDisData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("LevelDisData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	prof := ctx.QueryParam("prof")
	faygo.Infof("LevelDisData prof:[%v]", prof)

	leveldistributedata := make([]*level.LevelDisData, 0)
	err := db.GetLevelDisData(zeusid, stm, etm, utils.ParseInt(prof), &leveldistributedata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range leveldistributedata {
		pag.Push(v)
	}
	return ctx.JSONMsg(200, 200, pag)
})

/*
LevelLoseData
*/
var LevelLoseData = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	zeusid := ctx.QueryParam("zeusid")
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)
	faygo.Infof("LevelLoseData zeusid:[%s] timerange:[%v || %v]", zeusid, stm, etm)

	pag := &base.Pagination{PageSize: 20, PageIndex: 0, TotalItem: 20, TotalPage: 1, StartRow: 1, EndRow: 20}

	prof := ctx.QueryParam("prof")
	faygo.Infof("LevelLoseData prof:[%v]", prof)

	levellosedata := make([]*level.LevelLoseData, 0)
	err := db.GetLevelLoseData(zeusid, stm, etm, utils.ParseInt(prof), &levellosedata)
	if err != nil {
		pag.Warnings = fmt.Sprintf("%v", err)
	}

	for _, v := range levellosedata {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)

})
