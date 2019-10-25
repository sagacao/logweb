package pay

import "time"

//付费等级
type PayLevel struct {
	Level      int     `json:"level"`      //等级
	Amount     float32 `json:"amount"`     //付费金额
	PayerCount int64   `json:"payerCount"` //付费玩家数
	PayCount   int64   `json:"payCount"`   //付费次数
}

//付费小时
type PayHour struct {
	Hour       string  `json:"hour"`
	PayerCount int64   `json:"payerCount"`
	Amount     float32 `json:"amount"`
}

//每日付费
type PayDaily struct {
	DimDay     string  `json:"dimDay"`
	PayerCount int64   `json:"payerCount"`
	Amount     float32 `json:"amount"` //付费金额
}

//付费分析-用户贡献
type DaysPayAmount struct {
	DimDay        string  `json:"dimDay"`   //等级
	Dnp3          float32 `json:"dnp3"`     //
	PayCount3     int64   `json:"pay3"`     //3日付费总金额
	PlayerCount3  int64   `json:"player3"`  //3日活跃人数
	Dnp7          float32 `json:"dnp7"`     //
	PayCount7     int64   `json:"pay7"`     //7日付费总金额
	PlayerCount7  int64   `json:"player7"`  //7日活跃人数
	Dnp14         float32 `json:"dnp14"`    //
	PayCount14    int64   `json:"pay14"`    //14日付费总金额
	PlayerCount14 int64   `json:"player14"` //14日活跃人数
	Dnp30         float32 `json:"dnp30"`    //
	PayCount30    int64   `json:"pay30"`    //30日付费总金额
	PlayerCount30 int64   `json:"player30"` //30日活跃人数
	Dnp60         float32 `json:"dnp60"`    //
	PayCount60    int64   `json:"pay60"`    //60日付费总金额
	PlayerCount60 int64   `json:"player60"` //60日活跃人数

	Dnp90         float32 `json:"dnp90"`    //
	PayCount90    int64   `json:"pay90"`    //90日付费总金额
	PlayerCount90 int64   `json:"player90"` //90日活跃人数

	Dnp120         float32 `json:"dnp120"`    //
	PayCount120    int64   `json:"pay120"`    //120日付费总金额
	PlayerCount120 int64   `json:"player120"` //120日活跃人数

	Dnp150         float32 `json:"dnp150"`    //
	PayCount150    int64   `json:"pay150"`    //150日付费总金额
	PlayerCount150 int64   `json:"player150"` //150日活跃人数

	Dnp180         float32 `json:"dnp180"`    //
	PayCount180    int64   `json:"pay180"`    //180日付费总金额
	PlayerCount180 int64   `json:"player180"` //180日活跃人数
}

func InitDaysPayAmount(stm, etm time.Time, result *[]*DaysPayAmount) {
	tmcursor := stm
	for tmcursor.Before(etm) {
		*result = append(*result, &DaysPayAmount{
			DimDay:         tmcursor.Format("2006-01-02"),
			Dnp3:           0,
			PayCount3:      0,
			PlayerCount3:   0,
			Dnp7:           0,
			PayCount7:      0,
			PlayerCount7:   0,
			Dnp14:          0,
			PayCount14:     0,
			PlayerCount14:  0,
			Dnp30:          0,
			PayCount30:     0,
			PlayerCount30:  0,
			Dnp60:          0,
			PayCount60:     0,
			PlayerCount60:  0,
			Dnp90:          0,
			PayCount90:     0,
			PlayerCount90:  0,
			Dnp120:         0,
			PayCount120:    0,
			PlayerCount120: 0,
			Dnp150:         0,
			PayCount150:    0,
			PlayerCount150: 0,
			Dnp180:         0,
			PayCount180:    0,
			PlayerCount180: 0,
		})
		tmcursor = tmcursor.Add(time.Hour * 24)
	}
}

//每日流失
type PayLoss struct {
	DimDay    string `json:"dimDay"`
	Num       int    `json:"num"` //统计日
	LossNum3  int    `json:"lossNum3"`
	LossNum7  int    `json:"lossNum7"`
	LossNum14 int    `json:"lossNum14"`
	Dlr3      string `json:"dlr3"`
	Dlr7      string `json:"dlr7"`
	Dlr14     string `json:"dlr14"`
}

//渠道付费
type PayChannel struct {
	Channelid  int     `json:"channelid"`
	PayerCount int64   `json:"payerCount"`
	Amount     float32 `json:"amount"` //付费金额
}

//付费排行
type PayRank struct {
	Rank          int     `json:"rank"` //排名
	PlayerId      int     `json:"playerId"`
	Name          string  `json:"name"`   //昵称
	Level         int     `json:"level"`  //等级
	Amount        float32 `json:"amount"` //充值金额
	Gold          int64   `json:"gold"`   //剩余钻石数量
	Zeusid        int     `json:"zeusid"`
	Channelid     string  `json:"channelid"`
	CreateTime    string  `json:"createTime"`    //注册时间
	LastLoginTime string  `json:"lastLoginTime"` //最后登录时间
}

type RechargePrice struct {
	ChargeID   int   `json:"chargeId"`
	Price      int   `json:"price"`
	PayerCount int64 `json:"payerCount"`
	PayCount   int64 `json:"payCount"`
}
