package log

import "time"

type RoleIn struct {
	UserId   string `json:"userId" bson:"userid"`
	PlayerId string `json:"playerId" bson:"roleid"`
	LogTime  string `json:"logTime" bson:"logtime"`
}

type RoleOut struct {
	UserId     string `json:"userId" bson:"userid"`
	PlayerId   string `json:"playerId" bson:"roleid"`
	OnlineTime string `json:"onlineTime" bson:"time"`
	LogTime    string `json:"logTime" bson:"logtime"`
}

type RoleUp struct {
	UserId             string `json:"userId" bson:"userid"`
	PlayerId           string `json:"playerId" bson:"roleid"`
	Level              string `json:"level" bson:"lev"`
	Money              string `json:"money" bson:"totalcash"`
	BindGold           string `json:"bindGold" bson:"totalcash"`
	Gold               string `json:"gold" bson:"totalcash"`
	TotalPlayerTimeStr string `json:"totalPlayerTimeStr" bson:"lev"`
	LogTime            string `json:"logTime" bson:"logtime"`
}

///////////////////////////////////////////////////////////////
type RetRoleIn struct {
	UserId   string `json:"userid"`
	PlayerId string `json:"roleid"`
	LogTime  string `json:"logtime"`
}

type RetRoleOut struct {
	UserId     string `json:"userid"`
	PlayerId   string `json:"roleid"`
	OnlineTime string `json:"time"`
	LogTime    string `json:"logtime"`
}

type RetRoleUp struct {
	UserId             string `json:"userid"`
	PlayerId           string `json:"roleid"`
	Level              string `json:"lev"`
	Money              string `json:"totalcash"`
	BindGold           string `json:"-"`
	Gold               string `json:"-"`
	TotalPlayerTimeStr string `json:"-"`
	LogTime            string `json:"logtime"`
}

func PackagingRoleInQuery(page, pagecount int, stm, etm time.Time, roleid string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "logtime", "roleid"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"roleid": map[string]interface{}{
							"query":    roleid,
							"operator": "and",
						},
					},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"@timestamp": map[string]interface{}{
							"gt": stm.Format(time.RFC3339),
						},
					},
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

func PackagingRoleOutQuery(page, pagecount int, stm, etm time.Time, roleid string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "roleid", "time", "logtime"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"roleid": map[string]interface{}{
							"query":    roleid,
							"operator": "and",
						},
					},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"@timestamp": map[string]interface{}{
							"gt": stm.Format(time.RFC3339),
						},
					},
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

func PackagingRoleUpQuery(page, pagecount int, stm, etm time.Time, roleid string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "roleid", "logtime", "lev", "totalcash"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"roleid": map[string]interface{}{
							"query":    roleid,
							"operator": "and",
						},
					},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"@timestamp": map[string]interface{}{
							"gt": stm.Format(time.RFC3339),
						},
					},
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
