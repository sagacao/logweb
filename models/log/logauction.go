package log

import "time"

type PropsOn struct {
	AuctionId   string `json:"auctionId" bson:"auctionid"`
	UserId      string `json:"sellerId" bson:"roleid"`
	ItemId      string `json:"itemId" bson:"itemid"`
	ItemCount   string `json:"itemCount" bson:"itemnum"`
	Price       string `json:"price" bson:"price"`
	AuctionType string `json:"auctionType" bson:""`
	LogTime     string `json:"logTime" bson:"logtime"`
}

type PropsOut struct {
	AuctionId string `json:"auctionId" bson:"auctionid"`
	UserId    string `json:"sellerId" bson:"roleid"`
	ItemId    string `json:"itemId" bson:"itemid"`
	ItemCount string `json:"itemCount" bson:"itemnum"`
	Price     string `json:"price" bson:"price"`
	LogTime   string `json:"logTime" bson:"logtime"`
}

type PropsDone struct {
	AuctionId   string `json:"auctionId" bson:"auctionid"`
	UserId      string `json:"sellerId" bson:"buyerid"`
	BuyerId     string `json:"buyerId" bson:"roleid"`
	ItemId      string `json:"itemId"bson:"itemid"`
	ItemCount   string `json:"itemCount" bson:"itemnum"`
	Price       string `json:"price" bson:"price"`
	Fee         string `json:"fee" bson:"fee"`
	AuctionType string `json:"auctionType" bson:""`
	LogTime     string `json:"logTime" bson:"logtime"`
}

///////////////////////////////////////
type RetPropsBase struct {
	AuctionId   string `json:"auctionid"`
	UserId      string `json:"roleid"`
	ItemId      string `json:"itemid"`
	ItemCount   string `json:"itemnum"`
	Price       string `json:"price"`
	AuctionType string `json:"-"`
	LogTime     string `json:"logtime"`
}

type RetAuctionDone struct {
	AuctionId string `json:"auctionid"`
	BuyerId   string `json:"buyerid"`
	UserId    string `json:"roleid"`
	ItemId    string `json:"itemid"`
	ItemCount string `json:"itemnum"`
	Price     string `json:"price"`
	Fee       string `json:"fee"`
	LogTime   string `json:"logtime"`
}

func PackagingPropsBaseQuery(page, pagecount int, stm, etm time.Time, roleid, itemid string) map[string]interface{} {
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
		"_source": []string{"userid", "roleid", "logtime", "auctionid", "itemid", "itemnum", "starttime", "price"},
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

func PackagingPropsDoneQuery(page, pagecount int, stm, etm time.Time, roleid, itemid string) map[string]interface{} {
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
		"_source": []string{"roleid", "logtime", "auctionid", "itemid", "itemnum", "starttime", "price", "soldnum", "totalnum", "ownerid", "buyerid"},
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
