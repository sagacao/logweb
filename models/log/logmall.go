package log

import "time"

type MallSale struct {
	MallId    string `json:"mallId" bson:"mallid"`
	UserId    string `json:"userId" bson:"userid"`
	PlayerId  string `json:"playerId" bson:"roleid"`
	ItemId    int64  `json:"itemId" bson:"itemid"`
	ItemCount string `json:"itemCount" bson:"itemcount"`
	MoneyType string `json:"moneyType" bson:"yuanbaotype"`
	MoneyNeed string `json:"moneyNeed" bson:"yuanbao"`
	MoneyLeft uint32 `json:"moneyLeft"`
	LogTime   string `json:"logTime" bson:"logtime"`
}

type RetMallSale struct {
	MallId    string `json:"mallid"`
	UserId    string `json:"userid"`
	PlayerId  string `json:"roleid"`
	ItemId    string `json:"itemid"`
	ItemCount string `json:"itemcount"`
	MoneyType string `json:"yuanbaotype"`
	MoneyNeed string `json:"yuanbao"`
	LogTime   string `json:"logtime"`
}

func PackagingMallSaleQuery(page, pagecount int, stm, etm time.Time, mallid, roleid, itemid string) map[string]interface{} {
	filter := make([]map[string]interface{}, 0)
	filter = append(filter, map[string]interface{}{
		"range": map[string]interface{}{
			"@timestamp": map[string]interface{}{
				"gt": stm.Format(time.RFC3339),
			},
		},
	})

	if len(mallid) > 0 {
		filter = append(filter, map[string]interface{}{
			"term": map[string]interface{}{"mallid": mallid},
		})
	}

	if len(roleid) > 0 {
		filter = append(filter, map[string]interface{}{
			"term": map[string]interface{}{"roleid": roleid},
		})
	}

	if len(itemid) > 0 {
		filter = append(filter, map[string]interface{}{
			"term": map[string]interface{}{"itemid": itemid},
		})
	}

	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "roleid", "logtime", "mallid", "itemid", "itemcount", "yuanbaotype", "yuanbao"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{"filter": filter},
		},
		"sort": map[string]interface{}{
			"@timestamp": map[string]interface{}{
				"order": "asc",
			},
		},
	}
	return query
}
