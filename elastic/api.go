package elastic

import (
	"encoding/json"
	"fmt"
	"logweb/models/log"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

//订单查询
func GetAddGoldLogPage(zoneid string, logtype int, logvalue string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.ChargeOrder) (uint64, error) {
	searchindex := []string{fmt.Sprintf("charge_%s-*", zoneid)}
	searchtype := []string{"charge"}

	var query map[string]interface{}
	if logtype == 1 {
		query = log.PackagingChargeOrderByNickQuery(page, pagecount, starttime, endtime, logvalue)
	} else if logtype == 2 {
		query = log.PackagingChargeOrderByRoleQuery(page, pagecount, starttime, endtime, logvalue)
	} else {
		query = log.PackagingChargeOrderByOrderQuery(page, pagecount, starttime, endtime, logvalue)
	}
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetAddGoldLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetChargeOrder
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.ChargeOrder{
			ChannelOrderId: info.ChannelOrderId,
			UserId:         info.UserId,
			PlayerId:       info.PlayerId,
			PlayerName:     info.PlayerName,
			Level:          info.Level,
			Recharge:       info.Recharge,
			TotalRecharge:  info.TotalRecharge,
			CurGold:        info.CurGold,
			Channelid:      info.Channelid,
			LogTime:        info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

//账号登入
func GetAccountLoginLogPage(zoneid string, account string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.AccountIn) (uint64, error) {
	index := "accountlogin_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("accountlogin_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"accountlogin"}

	query := log.PackagingAccountInQuery(page, pagecount, starttime, endtime, account)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetAccountLoginLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetAccountIn
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.AccountIn{
			UserId:   info.UserId,
			Ip:       info.Ip,
			DeviceId: info.DeviceId,
			Device:   info.Device,
			LogTime:  info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

//账号登出
func GetAccountLogoutLogPage(zoneid string, account string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.AccountOut) (uint64, error) {
	index := "accountlogout_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("accountlogout_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"accountlogout"}

	query := log.PackagingAccountOutQuery(page, pagecount, starttime, endtime, account)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetAccountLogoutLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetAccountOut
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.AccountOut{
			UserId:   info.UserId,
			Ip:       info.Ip,
			DeviceId: info.DeviceId,
			LogTime:  info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

///////////////////////////////////////////////////////////////////////
/////////////// role
func GetPlayerLoginLogPage(zoneid string, roleid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.RoleIn) (uint64, error) {
	index := "rolelogin_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("rolelogin_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"rolelogin"}

	query := log.PackagingRoleInQuery(page, pagecount, starttime, endtime, roleid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetPlayerLoginLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetRoleIn
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.RoleIn{
			UserId:   info.UserId,
			PlayerId: info.PlayerId,
			LogTime:  info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

func GetPlayerLogoutLogPage(zoneid string, roleid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.RoleOut) (uint64, error) {
	index := "rolelogout_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("rolelogout_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"rolelogout"}

	query := log.PackagingRoleOutQuery(page, pagecount, starttime, endtime, roleid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetAccountLogoutLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetRoleOut
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.RoleOut{
			UserId:     info.UserId,
			PlayerId:   info.PlayerId,
			OnlineTime: info.OnlineTime,
			LogTime:    info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

func GetPlayerLevelUpLogPage(zoneid string, roleid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.RoleUp) (uint64, error) {
	index := "levelup_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("levelup_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"levelup"}

	query := log.PackagingRoleUpQuery(page, pagecount, starttime, endtime, roleid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetPlayerLevelUpLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetRoleUp
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.RoleUp{
			UserId:             info.UserId,
			PlayerId:           info.PlayerId,
			Level:              info.Level,
			Money:              info.Money,
			BindGold:           info.BindGold,
			Gold:               info.Gold,
			TotalPlayerTimeStr: info.TotalPlayerTimeStr,
			LogTime:            info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

///////////////////////////////////////////////////////////////////////
/////////////// mall
func GetPlayerDoMallLogPage(zoneid string, mallid, roleid, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.MallSale) (uint64, error) {
	index := "shoptrade_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("shoptrade_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"shoptrade"}

	query := log.PackagingMallSaleQuery(page, pagecount, starttime, endtime, mallid, roleid, itemid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetPlayerDoMallLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetMallSale
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.MallSale{
			MallId:    info.MallId,
			UserId:    info.UserId,
			PlayerId:  info.PlayerId,
			ItemId:    utils.ParseInt64(info.ItemId),
			ItemCount: info.ItemCount,
			MoneyType: info.MoneyType,
			MoneyNeed: info.MoneyNeed,
			MoneyLeft: 0,
			LogTime:   info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

///////////////////////////////////////////////////////////////////////
/////////////// auction
func GetAddAuctionLogPage(zoneid string, roleid, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.PropsOn) (uint64, error) {
	index := "auctionadd_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("auctionadd_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"auctionadd"}

	query := log.PackagingPropsBaseQuery(page, pagecount, starttime, endtime, roleid, itemid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetAddAuctionLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetPropsBase
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.PropsOn{
			AuctionId: info.AuctionId,
			UserId:    info.UserId,
			ItemId:    info.ItemId,
			ItemCount: info.ItemCount,
			Price:     info.Price,
			LogTime:   info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

func GetCancelAuctionLogPage(zoneid string, roleid, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.PropsOut) (uint64, error) {
	index := "auctiondel_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("auctiondel_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"auctiondel"}

	query := log.PackagingPropsBaseQuery(page, pagecount, starttime, endtime, roleid, itemid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetCancelAuctionLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetPropsBase
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.PropsOut{
			AuctionId: info.AuctionId,
			UserId:    info.UserId,
			ItemId:    info.ItemId,
			ItemCount: info.ItemCount,
			Price:     info.Price,
			LogTime:   info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

func GetDoneAuctionLogPage(zoneid string, roleid, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.PropsDone) (uint64, error) {
	index := "auctiondel_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("auctiondel_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"auctiondel"}

	query := log.PackagingPropsDoneQuery(page, pagecount, starttime, endtime, roleid, itemid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetCancelAuctionLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetAuctionDone
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.PropsDone{
			AuctionId: info.AuctionId,
			BuyerId:   info.BuyerId,
			UserId:    info.UserId,
			ItemId:    info.ItemId,
			ItemCount: info.ItemCount,
			Price:     info.Price,
			LogTime:   info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

///////////////////////////////////////////////////////////////////////
/////////////// item
func GetItemObtainLogPage(zoneid string, roleid string, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.ItemGet) (uint64, error) {
	index := "gainitem_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("gainitem_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"gainitem"}

	query := log.PackagingItemQuery(page, pagecount, starttime, endtime, roleid, itemid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetItemObtainLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetItemGet
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.ItemGet{
			PlayerId:  info.PlayerId,
			ItemId:    info.ItemId,
			Count:     info.Count,
			ChangeNum: info.ChangeNum,
			TypeName:  info.Reason,
			LogTime:   info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

func GetItemDisappearLogPage(zoneid string, roleid string, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.ItemLose) (uint64, error) {
	index := "loseitem_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("loseitem_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"loseitem"}

	query := log.PackagingItemQuery(page, pagecount, starttime, endtime, roleid, itemid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetItemDisappearLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetItemLose
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.ItemLose{
			PlayerId:  info.PlayerId,
			ItemId:    info.ItemId,
			Count:     info.Count,
			ChangeNum: info.ChangeNum,
			TypeName:  info.Reason,
			LogTime:   info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

///////////////////////////////////////////////////////////////////////
/////////////// task
func GetTaskAcceptLogPage(zoneid string, roleid string, taskid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.TaskAccept) (uint64, error) {
	index := "starttask_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("starttask_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"starttask"}

	query := log.PackagingTaskQuery(page, pagecount, starttime, endtime, roleid, taskid)
	dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetTaskAcceptLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetTaskAccept
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.TaskAccept{
			UserId:   info.TaskId,
			PlayerId: info.PlayerId,
			Level:    info.Level,
			Prof:     info.Prof,
			TaskId:   info.TaskId,
			LogTime:  info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

func GetTaskSuccessLogPage(zoneid string, roleid string, taskid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.TaskComplete) (uint64, error) {
	index := "endtask_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("endtask_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"endtask"}

	query := log.PackagingTaskQuery(page, pagecount, starttime, endtime, roleid, taskid)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetTaskSuccessLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetTaskComplete
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.TaskComplete{
			UserId:   info.TaskId,
			PlayerId: info.PlayerId,
			Level:    info.Level,
			Prof:     info.Prof,
			TaskId:   info.TaskId,
			LogTime:  info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

//货币产出消耗
func GetMoneyChangeLogPage(zoneid string, logtype int, roleId, moneytype string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.Money) (uint64, error) {
	index := "addyuanbao_*"
	zone := utils.ParseInt(zoneid)

	searchindex := []string{index}
	searchtype := make([]string, 0)
	if logtype == 1 {
		searchtype = append(searchtype, "addyuanbao")
		if zone != 0 {
			index = fmt.Sprintf("addyuanbao_%s-*", zone)
		}
	} else {
		searchtype = append(searchtype, "costyuanbao")
		if zone != 0 {
			index = fmt.Sprintf("costyuanbao_%s-*", zone)
		}
	}

	query := log.PackagingMoneyQuery(page, pagecount, starttime, endtime, roleId, moneytype)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetMoneyChangeLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetMoney
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.Money{
			PlayerId:  info.PlayerId,
			MoneyType: info.MoneyType,
			Count:     info.Count,
			Delta:     info.Delta,
			Reason:    info.Reason,
			LogTime:   info.LogTime,
		})
	}

	return rsp.Replys.Total, nil
}

//原始日志
func GetRawLogPage(zoneid string, logtype string, keywords string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.RawMsg) (uint64, error) {
	index := "raw_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("raw_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{"user"}

	query := log.PackagingRawMsgQuery(page, pagecount, starttime, endtime, keywords, logtype)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetRawLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		var info log.RetRawMsg
		err = utils.MapToStruct(v.Source, &info)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.RawMsg{
			LogTime:  info.LogTime,
			LogIndex: info.LogIndex,
			LogLevel: info.LogLevel,
			Message:  info.Message,
		})
	}

	return rsp.Replys.Total, nil
}

//{"userid": { "query": "1111"}, "roleid": { "query": "101"}}
func GetRawFormatLogPage(zoneid string, logtype string, keywords string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]*log.RawFormatMsg) (uint64, error) {
	index := fmt.Sprintf("%s_*", logtype) // "format_*"
	zone := utils.ParseInt(zoneid)
	if zone != 0 {
		index = fmt.Sprintf("%s_%v-*", logtype, zone) //fmt.Sprintf("format_%s-*", zone)
	}
	searchindex := []string{index}
	searchtype := []string{logtype}
	// faygo.Debugf("========== GetRawFormatLogPage ============")
	// faygo.Debugf("logtype:%s key:%s", logtype, keywords)
	// faygo.Debugf("========== GetRawFormatLogPage ============")

	query := log.PackagingRawFormatMsgQuery(page, pagecount, starttime, endtime, keywords, logtype)
	//dumpQueryByJson(query)

	rsp, err := elastic.searchHits(searchindex, searchtype, query)
	if rsp == nil {
		faygo.Warningf("GetRawFormatLogPage search --> %v", err)
		return 0, err
	}

	for _, v := range rsp.Replys.Hits {
		outrsp, err := json.Marshal(v.Source)
		if err != nil {
			faygo.Warningf("err %v ", err)
			continue
		}
		*results = append(*results, &log.RawFormatMsg{Message: string(outrsp)})
	}

	return rsp.Replys.Total, nil
}
