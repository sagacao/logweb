package log

import "time"

type AccountIn struct {
	UserId   string `json:"userId" bson:"userid"`
	Ip       string `json:"ip" bson:"loginip"`
	DeviceId string `json:"deviceId" bson:"macid"`
	Device   string `json:"device" bson:"model"`
	LogTime  string `json:"logTime" bson:"logtime"`
}

type AccountOut struct {
	UserId   string `json:"userId" bson:"userid`
	Ip       string `json:"ip" bson:"loginip`
	DeviceId string `json:"deviceId" bson:"macid"`
	LogTime  string `json:"logTime" bson:"logtime"`
}

/////////////////////////////////////////////////////
type RetAccountIn struct {
	UserId   string `json:"userid"`
	Ip       string `json:"loginip"`
	DeviceId string `json:"macid"`
	Device   string `json:"model"`
	LogTime  string `json:"logtime"`
}

type RetAccountOut struct {
	UserId   string `json:"userid"`
	Ip       string `json:"loginip"`
	DeviceId string `json:"macid"`
	LogTime  string `json:"logtime"`
}

func PackagingAccountInQuery(page, pagecount int, stm, etm time.Time, account string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "loginip", "logtime", "macid", "model"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"userid": map[string]interface{}{
							"query":    account,
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

func PackagingAccountOutQuery(page, pagecount int, stm, etm time.Time, account string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "loginip", "logtime", "macid"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"userid": map[string]interface{}{
							"query":    account,
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
