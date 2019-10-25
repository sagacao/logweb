package handler

import (
	"logweb/elastic"
	"logweb/models/base"
	"logweb/models/log"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

//订单查询
var GetAddGoldLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetAddGoldLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}

	nick := ctx.QueryParam("playerName")
	roleid := ctx.QueryParam("userId")
	roderid := ctx.QueryParam("channelOrderId")

	logtype := 0
	logvalue := roderid
	if len(nick) != 0 {
		logtype = 1
		logvalue = nick
	}
	if len(roleid) != 0 {
		logtype = 2
		logvalue = roleid
	}
	replylist := make([]*log.ChargeOrder, 0)
	total, err := elastic.GetAddGoldLogPage("5", logtype, logvalue, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)

	return ctx.JSONMsg(200, 200, pag)
})

//账号登入
var GetAccountLoginLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("getAccountLoginLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 10, StartRow: 0, EndRow: 1}
	userId := ctx.QueryParam("userId")
	if len(userId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}
	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.AccountIn, 0)

	total, err := elastic.GetAccountLoginLogPage(zeusid, userId, pageIndex, pagesize, stm, etm, &replylist)

	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}

	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)

	return ctx.JSONMsg(200, 200, pag)
})

//账号登出
var GetAccountLogoutLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("getAccountLogoutLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	userId := ctx.QueryParam("userId")
	if len(userId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.AccountOut, 0)
	total, err := elastic.GetAccountLogoutLogPage(zeusid, userId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)

	return ctx.JSONMsg(200, 200, pag)
})

//角色日志
var GetPlayerLoginLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetPlayerLoginLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	userId := ctx.QueryParam("userId")
	playerId := ctx.QueryParam("playerId")
	if len(userId) == 0 && len(playerId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.RoleIn, 0)
	total, err := elastic.GetPlayerLoginLogPage(zeusid, playerId, pageIndex, pagesize, stm, etm, &replylist)

	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}

	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)

	return ctx.JSONMsg(200, 200, pag)
})

var GetPlayerLogoutLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetPlayerLogoutLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	userId := ctx.QueryParam("userId")
	playerId := ctx.QueryParam("playerId")
	if len(userId) == 0 && len(playerId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.RoleOut, 0)
	total, err := elastic.GetPlayerLogoutLogPage(zeusid, playerId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)

	return ctx.JSONMsg(200, 200, pag)
})

var GetPlayerLevelUpLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetPlayerLevelUpLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	userId := ctx.QueryParam("userId")
	playerId := ctx.QueryParam("playerId")
	if len(userId) == 0 && len(playerId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.RoleUp, 0)
	total, err := elastic.GetPlayerLevelUpLogPage(zeusid, playerId, pageIndex, pagesize, stm, etm, &replylist)

	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)

	return ctx.JSONMsg(200, 200, pag)
})

///////////////////////////////////////////////////////////////////
////////////// Mall
var GetPlayerDoMallLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetPlayerDoMallLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 10, TotalPage: 1, StartRow: 0, EndRow: 10}
	roleId := ctx.QueryParam("roleId")
	mallId := ctx.QueryParam("mallId")
	itemId := ctx.QueryParam("itemId")
	if len(roleId) == 0 && len(mallId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.MallSale, 0)
	total, err := elastic.GetPlayerDoMallLogPage(zeusid, mallId, roleId, itemId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

///////////////////////////////////////////////////////////////////////////////
//////////////////// auction 拍卖行日志
var GetAddAuctionLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetAddAuctionLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 10, TotalPage: 1, StartRow: 0, EndRow: 10}

	sellerId := ctx.QueryParam("sellerId")
	itemId := ctx.QueryParam("itemId")
	if len(sellerId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.PropsOn, 0)
	total, err := elastic.GetAddAuctionLogPage(zeusid, sellerId, itemId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

//道具下架
var GetCancelAuctionLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetCancelAuctionLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 10, TotalPage: 1, StartRow: 0, EndRow: 10}

	sellerId := ctx.QueryParam("sellerId")
	itemId := ctx.QueryParam("itemId")
	if len(sellerId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.PropsOut, 0)
	total, err := elastic.GetCancelAuctionLogPage(zeusid, sellerId, itemId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

//道具成交
var GetDoneAuctionLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetDoneAuctionLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 10, TotalPage: 1, StartRow: 0, EndRow: 10}

	buyerId := ctx.QueryParam("buyerId")
	itemId := ctx.QueryParam("itemId")
	if len(buyerId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.PropsDone, 0)
	total, err := elastic.GetDoneAuctionLogPage(zeusid, buyerId, itemId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)

	return ctx.JSONMsg(200, 200, pag)
})

///////////////////////////////////////////////////////////////////////////////
//////////////////// task
var GetTaskAcceptLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetTaskAcceptLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}

	roleId := ctx.QueryParam("roleId")
	taskId := ctx.QueryParam("taskId")
	if len(roleId) == 0 && len(taskId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.TaskAccept, 0)
	total, err := elastic.GetTaskAcceptLogPage(zeusid, roleId, taskId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

var GetTaskSuccessLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetTaskSuccessLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	taskId := ctx.QueryParam("taskId")
	if len(roleId) == 0 && len(taskId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.TaskComplete, 0)
	total, err := elastic.GetTaskSuccessLogPage(zeusid, roleId, taskId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

///////////////////////////////////////////////////////////////////////////////
//////////////////// item 物品日志
//物品获取
var GetItemObtainLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetItemObtainLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	itemId := ctx.QueryParam("itemId")
	if len(roleId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.ItemGet, 0)
	total, err := elastic.GetItemObtainLogPage(zeusid, roleId, itemId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

var GetItemDisappearLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetItemDisappearLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	itemId := ctx.QueryParam("itemId")
	if len(roleId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.ItemLose, 0)
	total, err := elastic.GetItemDisappearLogPage(zeusid, roleId, itemId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

//物品合成（*）
var GetPlayerGemCompoundLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetPlayerGemCompoundLogPage")

	//stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	//etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))

	pag := &base.Pagination{PageSize: 1, PageIndex: 1, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	replylist := make([]*log.RoleUp, 0)
	//elastic.GetPlayerGemCompoundLogPage("5", "101", pageIndex, pagesize, stm, etm, &replylist)
	for _, v := range replylist {
		pag.Push(v)
	}

	return ctx.JSONMsg(200, 200, pag)
})

//货币产出消耗
var GetMoneyChangeLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetMoneyChangeLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	moneyType := ctx.QueryParam("moneyType")
	delta := ctx.QueryParam("delta")
	if len(roleId) == 0 || len(moneyType) == 0 || len(delta) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.Money, 0)
	total, err := elastic.GetMoneyChangeLogPage(zeusid, utils.ParseInt(delta), roleId, moneyType, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

//原始日志
var GetRawLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetRawLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	keywords := ctx.QueryParam("keywords")
	searchType := ctx.QueryParam("searchType")
	if len(keywords) == 0 || len(searchType) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.RawMsg, 0)
	total, err := elastic.GetRawLogPage(zeusid, searchType, keywords, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})

//分类日志
var GetRawFormatLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("GetRawFormatLogPage")

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}

	keywords := ctx.QueryParam("keywords") //"{\"userid\": { \"query\": \"1111\"}, \"roleid\": { \"query\": \"101\"}}"
	logType := ctx.QueryParam("logType")
	if len(keywords) == 0 || len(logType) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]*log.RawFormatMsg, 0)
	total, err := elastic.GetRawFormatLogPage(zeusid, logType, keywords, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)
	return ctx.JSONMsg(200, 200, pag)
})
