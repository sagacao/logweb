package log

type VirtualMoney struct {
	Time       string `json:"time" bson:"logtime"`
	CoinType   string `json:"coinType" bson:"type"`
	Count      string `json:"count" bson:"totalcharge"`
	TotalCount string `json:"totalCount" bson:"totalnum"`
	Channel    string `json:"channel" bson:"reason"`
	Depict     string `json:"depict" bson:"ext"`
}
