package mongo

import (
	"fmt"
	"logweb/models/log"
	"time"

	"github.com/henrylee2cn/faygo"

	"gopkg.in/mgo.v2/bson"
)

//账号登入
func GetAccountLoginLogPage(zoneid string, account string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.AccountIn) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}
	collection := mongosess.DB("format").C("accountlogin")

	query := bson.M{"userid": account, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.AccountIn
	for iter.Next(&rowdata) {
		// fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil

	// sql := func(c *mgo.Collection) error {
	// 	return c.Find(bson.M{"userid": account}).Limit(100).All(results)
	// }
	// err := mongoengine.witchCollection("format", "accountlogin", sql)
	// if err != nil {
	// 	return 0, nil
	// }
	// return 1, nil
}

//账号登出
func GetAccountLogoutLogPage(zoneid string, account string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.AccountOut) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"userid": account, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("accountlogout")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.AccountOut
	for iter.Next(&rowdata) {
		// fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//角色登入
func GetPlayerLoginLogPage(zoneid string, roleid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.RoleIn) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("rolelogin")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.RoleIn
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//角色登出
func GetPlayerLogoutLogPage(zoneid string, roleid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.RoleOut) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("rolelogout")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.RoleOut
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//角色升级
func GetPlayerLevelUpLogPage(zoneid string, roleid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.RoleUp) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("levelup")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.RoleUp
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//邮件发送日志
func GetMailSendLogPage(zoneid string, senderid, receiverid, mailid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.SendMail) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": senderid, "receiveid": receiverid, "gameuserid": mailid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("sendsysmail")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.SendMail
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//邮件删除日志
func GetMailDeleteLogPagee(zoneid string, receiverid, mailid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.DeleteMail) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": receiverid, "mailid": mailid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("deletemail")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.DeleteMail
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//商城日志
func GetPlayerDoMallLogPage(zoneid string, mallid, roleid, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.MallSale) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "mallid": mallid, "itemid": itemid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("shoptrade")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.MallSale
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//GetAddAuctionLogPage 拍卖行道具上架
func GetAddAuctionLogPage(zoneid string, roleid, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.PropsOn) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "itemid": itemid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("auctionadd")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.PropsOn
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//GetCancelAuctionLogPage 道具下架
func GetCancelAuctionLogPage(zoneid string, roleid, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.PropsOut) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "itemid": itemid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("auctiondel")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.PropsOut
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil

}

//GetDoneAuctionLogPage 道具成交
func GetDoneAuctionLogPage(zoneid string, roleid, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.PropsDone) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "itemid": itemid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("auctiondel")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.PropsDone
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//任务接取日志
func GetTaskAcceptLogPage(zoneid string, roleid string, taskid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.TaskAccept) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "taskid": taskid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("starttask")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.TaskAccept
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//任务完成日志
func GetTaskSuccessLogPage(zoneid string, roleid string, taskid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.TaskComplete) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "taskid": taskid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("endtask")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.TaskComplete
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//公会信息
func GetLeagueDataLogPage(zoneid string, guildId string, guildName string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.LeagueData) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": guildId, "leadername": guildName, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("leaguedata")
	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.LeagueData
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//物品获取
func GetItemObtainLogPage(zoneid string, roleid string, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.ItemGet) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	query := bson.M{"roleid": roleid, "itemid": itemid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}
	collection := mongosess.DB("format").C("gainitem")

	faygo.Infof("GetItemObtainLogPage", query, "page", page, "pagecount", pagecount)

	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}

	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.ItemGet
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//物品丢失
func GetItemDisappearLogPage(zoneid string, roleid string, itemid string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.ItemLose) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}
	collection := mongosess.DB("format").C("loseitem")

	query := bson.M{"roleid": roleid, "itemid": itemid, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}

	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.ItemLose
	for iter.Next(&rowdata) {
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//游戏币货币产出消耗
func GetMoneyChangeLogPage(zoneid string, logtype int, roleId, moneytype string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.Money) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	collection := mongosess.DB("format").C("costyuanbao")
	if logtype == 1 {
		collection = mongosess.DB("format").C("addyuanbao")
	}
	query := bson.M{"roleid": roleId, "type": moneytype, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}

	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.Money
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//原始日志
func GetRawLogPage(zoneid string, logtype string, keywords string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.RawMsg) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}
	collection := mongosess.DB("format").C("user")

	query := bson.M{"logtype": logtype, "keywords": keywords, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}

	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.RawMsg
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//分类日志
func GetRawFormatLogPage(zoneid string, logtype string, keywords string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.RawFormatMsg) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}
	collection := mongosess.DB("format").C("logtype")

	query := bson.M{"logtype": logtype, "keywords": keywords, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}

	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.RawFormatMsg
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

//虚拟货币日志
func GetVirtualMoneyLogPage(zoneid string, moneytype string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.VirtualMoney) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}
	collection := mongosess.DB("format").C("addyuanbao")

	query := bson.M{"type": moneytype, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}

	countNum, err := collection.Find(query).Count()
	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.VirtualMoney
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

// 侠客魂石
func GetVirtualPartnerSoulLogPage(zoneid string, logtype int, roleId string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.PartnerSoul) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	collection := mongosess.DB("format").C("partnersoulgain")
	query := bson.M{"roleid": roleId, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}

	countNum, err := collection.Find(query).Count()

	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.PartnerSoul
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}

// 侠客
func GetVirtualPartnerLogPage(zoneid string, logtype int, roleId string, page int, pagecount int, starttime time.Time, endtime time.Time, results *[]log.Partner) (uint64, error) {
	mongosess := mongoengine.getSession() //Session
	if mongosess == nil {
		return 0, nil
	}

	collection := mongosess.DB("format").C("partnergain")
	query := bson.M{"roleid": roleId, "logtime": bson.M{"$gte": starttime.Format(time.RFC3339), "$lt": endtime.Format(time.RFC3339)}}

	countNum, err := collection.Find(query).Count()

	if err != nil {
		return 0, fmt.Errorf("error : %v", err)
	}
	iter := collection.Find(query).Skip((page - 1) * pagecount).Limit(pagecount).Iter()
	var rowdata log.Partner
	for iter.Next(&rowdata) {
		fmt.Println(rowdata)
		*results = append(*results, rowdata)
	}
	return uint64(countNum), nil
}
