package league

type LeagueData struct {
	Number   string `json:"number"` //编号
	Level    string `json:"level"`  //等级
	Name     string `json:"name"`
	Camp     string `json:"camp"`
	Leader   string `json:"leader"`
	LeaderId string `json:"leaderid"`
	State    string `json:"state"`
	Brithday string `json:"brithday"`
}

type LeagueInfo struct {
	Logtime  string `json:"logtime"` //时间
	Level    int64 `json:"level"`   //等级
	Assets   int64 `json:"assets"`  //资产
	Credits  int64 `json:"credits"` //积分
	Health	 int64 `json:"health"`  //健康度
	Member	 int64 `json:"member"`  //成员
	Student	 int64 `json:"student"` //学徒
}