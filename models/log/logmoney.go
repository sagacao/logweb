package log

import "time"

type Money struct {
	PlayerId  string `json:"playerId" bson:"roleid"`
	MoneyType string `json:"moneyType" bson:"type"`
	Count     string `json:"totalCount" bson:"totalnum"`
	Delta     string `json:"delta" bson:"changecount"`
	LogTime   string `json:"logTime" bson:"logtime"`
	Reason    string `json:"reason" bson:"reason"`
	Depict    string `json:"depict" bson:"ext"`
}

type PartnerSoul struct {
	Lev      string `json:"lev" bson: "lev"`
	PlayerId string `json:"playerId" bson:"roleid"`
	Id       string `json:"id" bson: "id"`
	GainNum  string `json:"gainnum" bson: "gainnum"`
	TotalNum string `json:"totalnum" bson: "totalnum"`
	Reason   string `json:"reason" bson: "reason"`
	LogTime  string `json:"logTime" bson:"logtime"`
}

type Partner struct {
	Lev      string `json:"lev" bson: "lev"`
	PlayerId string `json:"playerId" bson:"roleid"`
	Id       string `json:"id" bson: "id"`
	GainNum  string `json:"gainnum" bson: "gainnum"`
	TotalNum string `json:"totalnum" bson: "totalnum"`
	Reason   string `json:"reason" bson: "reason"`
	LogTime  string `json:"logTime" bson:"logtime"`
}

type RetMoney struct {
	PlayerId  string `json:"roleid"`
	MoneyType string `json:"type"`
	Count     string `json:"totalcash"`
	Delta     string `json:"changecount"`
	LogTime   string `json:"logtime"`
	Reason    string `json:"reason"`
}

func PackagingMoneyQuery(page, pagecount int, stm, etm time.Time, roleid, moneytype string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"roleid", "logtime", "type", "totalcash", "changecount", "reason"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"range": map[string]interface{}{
							"@timestamp": map[string]interface{}{
								"gt": stm.Format(time.RFC3339),
								"lt": etm.Format(time.RFC3339),
							},
						},
					},
					{"term": map[string]interface{}{"type": moneytype}},
					{"term": map[string]interface{}{"roleid": roleid}},
				},
			},
		},
		"sort": map[string]interface{}{
			"@timestamp": map[string]interface{}{
				"order": "asc",
			},
		},
	}

	return query
}
