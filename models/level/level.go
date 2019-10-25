package level

type LevelDisData struct {
	Plat    int   `json:"plat"`		    //渠道
	Level   int64 `json:"level"` 		//等级
	Num     int64 `json:"num"`   		//人数
}

type LevelLoseData struct {
	Plat 	int   `json:"plat"`			//渠道
	Level   int64 `json:"level"` 		//等级
	Num     int64 `json:"num"`   		//人数
}