package handler

import (
	"logweb/models/base"
	"logweb/models/log"
	"logweb/mongo"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

//账号登入
var AccountLoginLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 10, StartRow: 0, EndRow: 1}
	userId := ctx.QueryParam("userId") //accountId
	faygo.Infof("getAccountLoginLogPage >>>>>>>>>>>>>> tm:[%v|%v] user:[%v] page[%v|%v]", stm, etm, userId, pageIndex, pagesize)
	if len(userId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}
	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.AccountIn, 0)
	total, err := mongo.GetAccountLoginLogPage(zeusid, userId, pageIndex, pagesize, stm, etm, &replylist)
	if err != nil {
		return ctx.JSONMsg(200, 200, pag)
	}
	faygo.Infof("return total size:[%v]", total)

	for _, v := range replylist {
		pag.Push(v)
	}

	pag.TotalItem = int(total)
	pag.TotalPage = int(pag.TotalItem / pag.PageSize)

	return ctx.JSONMsg(200, 200, pag)
})

//账号登出
var AccountLogoutLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	userId := ctx.QueryParam("userId")
	faygo.Infof("AccountLogoutLogPage >>>>>>>>>>>>>> tm:[%v|%v] user:[%v]", stm, etm, userId)
	if len(userId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.AccountOut, 0)
	total, err := mongo.GetAccountLogoutLogPage(zeusid, userId, pageIndex, pagesize, stm, etm, &replylist)
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

//角色登入日志
var PlayerLoginLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	playerId := ctx.QueryParam("playerId")
	faygo.Infof("PlayerLoginLogPage >>>>>>>>>>>>>> tm:[%v|%v] user:[%v]", stm, etm, playerId)
	if len(playerId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.RoleIn, 0)
	total, err := mongo.GetPlayerLoginLogPage(zeusid, playerId, pageIndex, pagesize, stm, etm, &replylist)
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

//角色登出日志
var PlayerLogoutLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	playerId := ctx.QueryParam("playerId")
	faygo.Infof("PlayerLogoutLogPage >>>>>>>>>>>>>> tm:[%v|%v] user:[%v]", stm, etm, playerId)
	if len(playerId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.RoleOut, 0)
	total, err := mongo.GetPlayerLogoutLogPage(zeusid, playerId, pageIndex, pagesize, stm, etm, &replylist)
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

//角色升级日志
var PlayerLevelUpLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	playerId := ctx.QueryParam("playerId")
	faygo.Infof("PlayerLevelUpLogPage >>>>>>>>>>>>>> tm:[%v|%v] playerId:[%v]", stm, etm, playerId)
	if len(playerId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.RoleUp, 0)
	total, err := mongo.GetPlayerLevelUpLogPage(zeusid, playerId, pageIndex, pagesize, stm, etm, &replylist)
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

//邮件发送日志
var MailSendLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	senderId := ctx.QueryParam("senderId")
	receiverId := ctx.QueryParam("receiverId")
	mailId := ctx.QueryParam("mailId")

	faygo.Infof("PlayerLevelUpLogPage >>>>>>>>>>>>>> tm:[%v|%v] senderId:[%v] receiverId:[%v] mailId:[%v]", stm, etm, senderId, receiverId, mailId)
	if len(senderId) == 0 && len(receiverId) == 0 && len(mailId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.SendMail, 0)
	total, err := mongo.GetMailSendLogPage(zeusid, senderId, receiverId, mailId, pageIndex, pagesize, stm, etm, &replylist)
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

//邮件删除日志
var MailDeleteLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	receiverId := ctx.QueryParam("receiverId")
	mailId := ctx.QueryParam("mailId")

	faygo.Infof("MailDeleteLogPage >>>>>>>>>>>>>> tm:[%v|%v] receiverId:[%v] mailId:[%v]", stm, etm, receiverId, mailId)
	if len(receiverId) == 0 && len(mailId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.DeleteMail, 0)
	total, err := mongo.GetMailDeleteLogPagee(zeusid, receiverId, mailId, pageIndex, pagesize, stm, etm, &replylist)
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

//商城日志
var PlayerDoMallLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 10, TotalPage: 1, StartRow: 0, EndRow: 10}
	roleId := ctx.QueryParam("roleId")
	mallId := ctx.QueryParam("mallId")
	itemId := ctx.QueryParam("itemId")
	faygo.Infof("PlayerDoMallLogPage >>>>>>>>>>>>>> tm:[%v|%v] roleId:[%v] mallId:[%v] itemId:[%v]", stm, etm, roleId, mallId, itemId)
	if len(roleId) == 0 && len(mallId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.MallSale, 0)
	total, err := mongo.GetPlayerDoMallLogPage(zeusid, mallId, roleId, itemId, pageIndex, pagesize, stm, etm, &replylist)
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

//拍卖行道具上架
var AddAuctionLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 10, TotalPage: 1, StartRow: 0, EndRow: 10}

	sellerId := ctx.QueryParam("sellerId")
	itemId := ctx.QueryParam("itemId")
	faygo.Infof("AddAuctionLogPage >>>>>>>>>>>>>> tm:[%v|%v] sellerId:[%v] itemId:[%v]", stm, etm, sellerId, itemId)
	if len(sellerId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.PropsOn, 0)
	total, err := mongo.GetAddAuctionLogPage(zeusid, sellerId, itemId, pageIndex, pagesize, stm, etm, &replylist)
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
var CancelAuctionLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 10, TotalPage: 1, StartRow: 0, EndRow: 10}

	sellerId := ctx.QueryParam("sellerId")
	itemId := ctx.QueryParam("itemId")
	faygo.Infof("CancelAuctionLogPage >>>>>>>>>>>>>> tm:[%v|%v] sellerId:[%v] itemId:[%v]", stm, etm, sellerId, itemId)
	if len(sellerId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.PropsOut, 0)
	total, err := mongo.GetCancelAuctionLogPage(zeusid, sellerId, itemId, pageIndex, pagesize, stm, etm, &replylist)
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
var DoneAuctionLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 10, TotalPage: 1, StartRow: 0, EndRow: 10}

	buyerId := ctx.QueryParam("buyerId")
	itemId := ctx.QueryParam("itemId")
	faygo.Infof("DoneAuctionLogPage >>>>>>>>>>>>>> tm:[%v|%v] buyerId:[%v] itemId:[%v]", stm, etm, buyerId, itemId)
	if len(buyerId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.PropsDone, 0)
	total, err := mongo.GetDoneAuctionLogPage(zeusid, buyerId, itemId, pageIndex, pagesize, stm, etm, &replylist)
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

//任务接取日志
var TaskAcceptLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}

	roleId := ctx.QueryParam("roleId")
	taskId := ctx.QueryParam("taskId")
	faygo.Infof("TaskAcceptLogPage >>>>>>>>>>>>>> tm:[%v|%v] roleId:[%v] taskId:[%v]", stm, etm, roleId, taskId)
	if len(roleId) == 0 && len(taskId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.TaskAccept, 0)
	total, err := mongo.GetTaskAcceptLogPage(zeusid, roleId, taskId, pageIndex, pagesize, stm, etm, &replylist)
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

//任务完成日志
var TaskSuccessLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	taskId := ctx.QueryParam("taskId")
	faygo.Infof("TaskSuccessLogPage >>>>>>>>>>>>>> tm:[%v|%v] roleId:[%v] taskId:[%v]", stm, etm, roleId, taskId)
	if len(roleId) == 0 && len(taskId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.TaskComplete, 0)
	total, err := mongo.GetTaskSuccessLogPage(zeusid, roleId, taskId, pageIndex, pagesize, stm, etm, &replylist)
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

//公会信息
var LeagueDataLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	guildId := ctx.QueryParam("guildId")
	guildName := ctx.QueryParam("guildName")
	faygo.Infof("LeagueDataLogPage >>>>>>>>>>>>>> tm:[%v|%v] guildId:[%v] guildName:[%v]", stm, etm, guildId, guildName)
	if len(guildId) == 0 && len(guildName) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.LeagueData, 0)
	total, err := mongo.GetLeagueDataLogPage(zeusid, guildId, guildName, pageIndex, pagesize, stm, etm, &replylist)
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

//物品获取
var ItemObtainLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	itemId := ctx.QueryParam("itemId")
	faygo.Infof("ItemObtainLogPage >>>>>>>>>>>>>> tm:[%v|%v] roleId:[%v] itemId:[%v]", stm, etm, roleId, itemId)
	if len(roleId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.ItemGet, 0)
	total, err := mongo.GetItemObtainLogPage(zeusid, roleId, itemId, pageIndex, pagesize, stm, etm, &replylist)
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

//物品丢失
var ItemDisappearLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	itemId := ctx.QueryParam("itemId")
	faygo.Infof("ItemDisappearLogPage >>>>>>>>>>>>>> tm:[%v|%v] roleId:[%v] itemId:[%v]", stm, etm, roleId, itemId)
	if len(roleId) == 0 && len(itemId) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.ItemLose, 0)
	total, err := mongo.GetItemDisappearLogPage(zeusid, roleId, itemId, pageIndex, pagesize, stm, etm, &replylist)
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

//游戏货币产出消耗
var MoneyChangeLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	moneyType := ctx.QueryParam("moneyType")
	delta := ctx.QueryParam("delta")
	faygo.Infof("MoneyChangeLogPage >>>>>>>>>>>>>> tm:[%v|%v] roleId:[%v] moneyType:[%v] delta:[%v]", stm, etm, roleId, moneyType, delta)
	if len(roleId) == 0 || len(moneyType) == 0 || len(delta) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.Money, 0)
	total, err := mongo.GetMoneyChangeLogPage(zeusid, utils.ParseInt(delta), roleId, moneyType, pageIndex, pagesize, stm, etm, &replylist)
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
var RawLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	keywords := ctx.QueryParam("keywords")
	searchType := ctx.QueryParam("searchType")
	faygo.Infof("RawLogPage >>>>>>>>>>>>>> tm:[%v|%v] keywords:[%v] searchType:[%v]", stm, etm, keywords, searchType)
	if len(keywords) == 0 || len(searchType) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.RawMsg, 0)
	total, err := mongo.GetRawLogPage(zeusid, searchType, keywords, pageIndex, pagesize, stm, etm, &replylist)
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
var RawFormatLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	keywords := ctx.QueryParam("keywords")
	logType := ctx.QueryParam("logType")
	faygo.Infof("RawFormatLogPage >>>>>>>>>>>>>> tm:[%v|%v] keywords:[%v] logType:[%v]", stm, etm, keywords, logType)
	if len(keywords) == 0 || len(logType) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.RawFormatMsg, 0)
	total, err := mongo.GetRawFormatLogPage(zeusid, logType, keywords, pageIndex, pagesize, stm, etm, &replylist)
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

//虚拟货币日志
var VirtualMoneyLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}

	Type := ctx.QueryParam("type")
	faygo.Infof("VirtualMoneyLogPage >>>>>>>>>>>>>> tm:[%v|%v] Type:[%v]", stm, etm, Type)
	if len(Type) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.VirtualMoney, 0)
	total, err := mongo.GetVirtualMoneyLogPage(zeusid, Type, pageIndex, pagesize, stm, etm, &replylist)
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

//侠客魂石
var VirtualPartnerSoulLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	delta := ctx.QueryParam("delta")
	faygo.Infof("VirtualPartnerSoulLogPage >>>>>>>>>>>>>> tm:[%v|%v] roleId:[%v] delta:[%v]", stm, etm, roleId, delta)
	if len(roleId) == 0 || len(delta) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.PartnerSoul, 0)
	total, err := mongo.GetVirtualPartnerSoulLogPage(zeusid, utils.ParseInt(delta), roleId, pageIndex, pagesize, stm, etm, &replylist)
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

// 侠客
var VirtualPartnerLogPage = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	stm, _ := time.Parse("2006-01-02", ctx.QueryParam("startTime"))
	etm, _ := time.Parse("2006-01-02", ctx.QueryParam("endTime"))
	etm = etm.Add(24 * time.Hour)

	pageIndex := utils.ParseInt(ctx.QueryParam("pageIndex"))
	pagesize := utils.ParseInt(ctx.QueryParam("pageSize"))
	pag := &base.Pagination{PageSize: pagesize, PageIndex: pageIndex, TotalItem: 1, TotalPage: 1, StartRow: 0, EndRow: 1}
	roleId := ctx.QueryParam("roleId")
	delta := ctx.QueryParam("delta")
	faygo.Infof("VirtualPartnerLogPage >>>>>>>>>>>>>> tm:[%v|%v] roleId:[%v] delta:[%v]", stm, etm, roleId, delta)
	if len(roleId) == 0 || len(delta) == 0 {
		return ctx.JSONMsg(200, 200, pag)
	}

	zeusid := ctx.QueryParam("zeusid")
	replylist := make([]log.Partner, 0)
	total, err := mongo.GetVirtualPartnerLogPage(zeusid, utils.ParseInt(delta), roleId, pageIndex, pagesize, stm, etm, &replylist)
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
