package log

import "time"

type ItemGet struct {
	PlayerId  string `json:"playerId" bson:"roleid"`
	ItemId    string `json:"itemId" bson:"itemid"`
	Count     string `json:"count" bson:"itemcount"`
	ChangeNum string `json:"changenum" bson:"changenum"`
	TypeName  string `json:"typeName" bson:"reason"`
	LogTime   string `json:"logTime" bson:"logtime"`
	Depict    string `json:"depict" bson:"ext"`
}

type ItemLose struct {
	PlayerId  string `json:"playerId" bson:"roleid"`
	ItemId    string `json:"itemId" bson:"itemid"`
	Count     string `json:"count" bson:"itemcount"`
	ChangeNum string `json:"changenum" bson:"changenum"`
	TypeName  string `json:"typeName" bson:"reason"`
	LogTime   string `json:"logTime" bson:"logtime"`
	Depict    string `json:"depict" bson:"ext"`
}

//////////////////////////////////////////////////////
type RetItemGet struct {
	PlayerId  string `json:"roleid"`
	ItemId    string `json:"itemid"`
	Count     string `json:"itemcount"`
	ChangeNum string `json:"changenum"`
	LogTime   string `json:"logtime"`
	Reason    string `json:"reason"`
	Ext       string `json:"ext"`
}

type RetItemLose struct {
	PlayerId  string `json:"roleid"`
	ItemId    string `json:"itemid"`
	Count     string `json:"itemcount"`
	ChangeNum string `json:"changenum"`
	LogTime   string `json:"logtime"`
	Reason    string `json:"reason"`
	Ext       string `json:"ext"`
}

func PackagingItemQuery(page, pagecount int, stm, etm time.Time, roleid, itemid string) map[string]interface{} {
	filter := make([]map[string]interface{}, 0)
	filter = append(filter, map[string]interface{}{
		"range": map[string]interface{}{
			"@timestamp": map[string]interface{}{
				"gt": stm.Format(time.RFC3339),
			},
		},
	})
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
		"_source": []string{"userid", "roleid", "logtime", "itemid", "itemcount", "changenum", "reason", "ext"},
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
