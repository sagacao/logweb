package cost

import (
	"time"
)

//货币消耗
type MoneyChange struct {
	Reason string `json:"reason"`
	Num    int64  `json:"num"`
}

//虚拟币产出消耗汇总
type MoneyChangeSummary struct {
	DimDay      string `json:"dimDay"`
	MoneyGain   int64  `json:"moneyGain"`
	MoneyCost   int64  `json:"moneyCost"`
	MoneyRemain int64  `json:"moneyRemain"`
	DimmiGain   int64  `json:"dimmiGain"`
	DimmiCost   int64  `json:"dimmiCost"`
	DimmiRemain int64  `json:"dimmiRemain"`
	GoldGain    int64  `json:"goldGain"`
	GoldCost    int64  `json:"goldCost"`
	GoldRemain  int64  `json:"goldRemain"`
}

func InitMoneyChangeSummary(stm, etm time.Time, result *[]*MoneyChangeSummary) {
	tmcursor := stm
	for tmcursor.Before(etm) {
		*result = append(*result, &MoneyChangeSummary{
			DimDay:      tmcursor.Format("2006-01-02"),
			MoneyGain:   0,
			MoneyCost:   0,
			MoneyRemain: 0,
			DimmiGain:   0,
			DimmiCost:   0,
			DimmiRemain: 0,
			GoldGain:    0,
			GoldCost:    0,
			GoldRemain:  0,
		})
		tmcursor = tmcursor.Add(time.Hour * 24)
	}
}

type Moneysummary struct {
	DimDay    string `json:"dimDay"`
	MoneyGain int64  `json:"moneyGain"`
	MoneyCost int64  `json:"moneyCost"`
}

type Dimmisummary struct {
	DimDay    string `json:"dimDay"`
	DimmiGain int64  `json:"dimmiGain"`
	DimmiCost int64  `json:"dimmiCost"`
}

type Goldsummary struct {
	DimDay   string `json:"dimDay"`
	GoldGain int64  `json:"goldGain"`
	GoldCost int64  `json:"goldCost"`
}

type RechargeSummary struct {
	ServerName string  `json:"servername"`
	Time       string  `json:"time"`
	Cash       float32 `json:"cash"`
}
