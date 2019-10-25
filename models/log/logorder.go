package log

import "time"

type ChargeOrder struct {
	ChannelOrderId string `json:"channelOrderId"`
	UserId         string `json:"userId"`
	PlayerId       string `json:"playerId"`
	PlayerName     int64  `json:"playerName"`
	Level          string `json:"level"`
	Recharge       string `json:"recharge"`
	TotalRecharge  string `json:"totalRecharge"`
	CurGold        uint32 `json:"curGold"`
	Channelid      uint32 `json:"channelid"`
	LogTime        string `json:"logTime"`
}

type RetChargeOrder struct {
	ChannelOrderId string `json:"gameorder"`
	UserId         string `json:"userid"`
	PlayerId       string `json:"roleid"`
	PlayerName     int64  `json:"account"`
	Level          string `json:"lev"`
	Recharge       string `json:"price"`
	TotalRecharge  string `json:"totalcash"`
	CurGold        uint32 `json:"-"`
	Channelid      uint32 `json:"platform"`
	LogTime        string `json:"logtime"`
}

func PackagingChargeOrderByNickQuery(page, pagecount int, stm, etm time.Time, value string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "roleid", "logtime", "account", "logtime", "platform", "lev", "price", "totalcash", "gameorder", "chargetype"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"range": map[string]interface{}{
							"@timestamp": map[string]interface{}{
								"gt": stm.Format(time.RFC3339),
							},
						},
					},
					{
						"terms": map[string]interface{}{
							"account": []string{value},
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

func PackagingChargeOrderByRoleQuery(page, pagecount int, stm, etm time.Time, value string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "roleid", "logtime", "account", "logtime", "platform", "lev", "price", "totalcash", "gameorder", "chargetype"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"range": map[string]interface{}{
							"@timestamp": map[string]interface{}{
								"gt": stm.Format(time.RFC3339),
							},
						},
					},
					{
						"terms": map[string]interface{}{
							"roleid": []string{value},
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

func PackagingChargeOrderByOrderQuery(page, pagecount int, stm, etm time.Time, value string) map[string]interface{} {
	query := map[string]interface{}{
		"from":    (page - 1) * pagecount,
		"size":    pagecount,
		"_source": []string{"userid", "roleid", "logtime", "account", "logtime", "platform", "lev", "price", "totalcash", "gameorder", "chargetype"},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"range": map[string]interface{}{
							"@timestamp": map[string]interface{}{
								"gt": stm.Format(time.RFC3339),
							},
						},
					},
					{
						"terms": map[string]interface{}{
							"gameorder": []string{value},
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
