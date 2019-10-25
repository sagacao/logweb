package log

import "time"

type TaskAccept struct {
	UserId   string `json:"userId" bson:"userid"`
	PlayerId string `json:"playerId" bson:"roleid"`
	Level    string `json:"level" bson:"lev"`
	Prof     string `json:"occupation" bson:"prof"`
	TaskId   string `json:"taskId" bson:"taskid"`
	LogTime  string `json:"logTime" bson:"logtime"`
}

type TaskComplete struct {
	UserId        string `json:"userId" bson:"userid"`
	PlayerId      string `json:"playerId" bson:"roleid"`
	Level         string `json:"level" bson:"lev"`
	Prof          string `json:"occupation" bson:"prof"`
	TaskId        string `json:"taskId" bson:"taskid"`
	Money         string `json:"money" bson:""`
	Exp           string `json:"exp" bson:""`
	FullRepuIndex string `json:"fullRepuIndex" bson:""`
	ItemIds       string `json:"itemIds" bson:""`
	LogTime       string `json:"logTime" bson:"logtime"`
}

//////////////////////////////////////////////////////
type RetTaskAccept struct {
	UserId   string `json:"userid"`
	PlayerId string `json:"roleid"`
	Level    string `json:"lev"`
	Prof     string `json:"prof"`
	TaskId   string `json:"taskid"`
	LogTime  string `json:"logtime"`
}

type RetTaskComplete struct {
	UserId        string `json:"userid"`
	PlayerId      string `json:"roleid"`
	Level         string `json:"lev"`
	Prof          string `json:"prof"`
	TaskId        string `json:"taskid"`
	Money         string `json:"-"`
	Exp           string `json:"-"`
	FullRepuIndex string `json:"-"`
	ItemIds       string `json:"-"`
	LogTime       string `json:"logtime"`
}

func PackagingTaskQuery(page, pagecount int, stm, etm time.Time, roleid, taskid string) map[string]interface{} {
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

	if len(taskid) > 0 {
		filter = append(filter, map[string]interface{}{
			"term": map[string]interface{}{"taskid": taskid},
		})
	}

	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "roleid", "logtime", "taskid", "lev", "prof"},
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
