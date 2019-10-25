package handler

import (
	"logweb/ini"

	"github.com/henrylee2cn/faygo"
)

/*
Config
*/
var Config = faygo.HandlerFunc(func(ctx *faygo.Context) error {
	faygo.Info("Index")

	index_data_map := make(map[string]interface{})

	// index_data_map["moneyChangeReasonMap"] = ini.Itemchange.Info
	// index_data_map["itemMallMap"] = ini.ItemMall.Info
	// index_data_map["mailTypeMap"] = ini.MailType.Info
	// index_data_map["channelMap"] = ini.Platform.Info
	index_data_map["zeusTypeMap"] = ini.ZeusType.Info
	index_data_map["platTypeMap"] = ini.Platform.Info
	index_data_map["moneyChangeReasonMap"] = ini.Itemchange.Info
	index_data_map["rechargeMap"] = ini.Rechargeform.Info
	
	return ctx.JSONMsg(200, 200, index_data_map)
})
