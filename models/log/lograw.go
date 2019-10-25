package log

import (
	"encoding/json"
	"time"

	"github.com/henrylee2cn/faygo"
)

type RawMsg struct {
	LogTime  string `json:"logTime" bson:"logtime"`
	LogIndex string `json:"logIndex" bson:"index"`
	LogLevel string `json:"LogLevel" bson:"level"`
	Message  string `json:"msg" bson:"msg"`
}

type RawFormatMsg struct {
	Message string `json:"msg" bson:"msg"`
}

////////////////////////////////////////////////////
type RetRawMsg struct {
	LogTime  string `json:"logtime"`
	LogIndex string `json:"index"`
	LogLevel string `json:"level"`
	Message  string `json:"msg"`
}

func PackagingRawMsgQuery(page, pagecount int, stm, etm time.Time, keywords, logtype string) map[string]interface{} {
	query := map[string]interface{}{
		"from": (page - 1) * pagecount,
		"size": pagecount,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"match": map[string]interface{}{
						"msg": map[string]interface{}{
							"query":    keywords,
							"operator": logtype,
						},
					},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"@timestamp": map[string]interface{}{
							"gt": stm.Format(time.RFC3339),
							"lt": etm.Format(time.RFC3339),
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

func PackagingRawFormatMsgQuery(page, pagecount int, stm, etm time.Time, keywords, logtype string) map[string]interface{} {
	var matchmap map[string]interface{}
	err := json.Unmarshal([]byte(keywords), &matchmap)
	if err != nil {
		faygo.Warningf("err %v ", err)
		return nil
	}

	must := make([]map[string]interface{}, 0)
	for k, v := range matchmap {
		must = append(must, map[string]interface{}{"match": map[string]interface{}{k: v}})
	}

	query := map[string]interface{}{
		"from": (page - 1) * pagecount,
		"size": pagecount,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": must,
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"@timestamp": map[string]interface{}{
							"gt": stm.Format(time.RFC3339),
							"lt": etm.Format(time.RFC3339),
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
