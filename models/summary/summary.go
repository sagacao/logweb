package summary

import "time"

type GameSummary struct {
	DimDay            string  `json:"dimDay"`
	TotalIncome       float32 `json:"totalIncome"`       //当日总收入（货币）
	RechargeUsers     int64   `json:"rechargeUsers"`     //当日充值人数
	Dnu               int64   `json:"dnu"`               //当日新增账号数
	MacDnu            int64   `json:"macDnu"`            //当日新增设备数
	NewPayRate        string  `json:"newPayRate"`        //新增当日付费率（百分数） ( NewRechargeUsers / Dnu)
	PayRate           string  `json:"payRate"`           //当日付费率（百分数）(RechargeUsers / Dau)
	Arppu             string  `json:"arppu"`             //当日付费玩家平均收入（货币） ( TotalIncome / RechargeUsers)
	NewRechargeIncome float32 `json:"newRechargeIncome"` //新增充值收入（货币）
	NewRechargeUsers  int64   `json:"newRechargeUsers"`  //新增充值用户
	Dau               int64   `json:"dau"`               //日活跃（登录过）玩家数
	MacDau            int64   `json:"macDau"`            //日活跃（登录过）设备数
	WithoutNewDau     int32   `json:"withoutNewDau"`     //去新日活跃玩家数(dau-dnu)
	Ndrr1             string  `json:"ndrr1"`             //次日留存率(百分数)
	Pcu               int32   `json:"pcu"`               //最高同时在线玩家数
	Acu               int32   `json:"acu"`               //平均同时在线玩家数
	AvgOnlineTime     float32 `json:"avgOnlineTime"`     //人均在线时长（min）（小数）
}

func InitGameSummary(stm, etm time.Time, result *[]*GameSummary) {
	tmcursor := stm
	for tmcursor.Before(etm) {
		*result = append(*result, &GameSummary{
			DimDay:            tmcursor.Format("2006-01-02"),
			TotalIncome:       0,
			RechargeUsers:     0,
			Dnu:               0,
			MacDnu:            0,
			NewPayRate:        "0%",
			PayRate:           "0%",
			Arppu:             "0",
			NewRechargeIncome: 0,
			NewRechargeUsers:  0,
			Dau:               0,
			MacDau:            0,
			WithoutNewDau:     0,
			Ndrr1:             "0",
			Pcu:               0,
			Acu:               0,
			AvgOnlineTime:     0,
		})
		tmcursor = tmcursor.Add(time.Hour * 24)
	}
}
