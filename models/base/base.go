package base

type Pagination struct {
	PageSize  int    `json:"pageSize"`
	PageIndex int    `json:"pageIndex"`
	TotalItem int    `json:"totalItem"`
	TotalPage int    `json:"totalPage"`
	StartRow  int    `json:"startRow"`
	EndRow    int    `json:"endRow"`
	Warnings  string `json:"warnings"`

	Rows []interface{} `json:"rows"`
}

func (p *Pagination) Push(row interface{}) {
	p.Rows = append(p.Rows, row)
}

type DayDnuData struct {
	DimDay  string `json:"dimDay"`
	Dnu     int64  `json:"dnu"`     //当日新增账号数
	RoleDnu int64  `json:"roleDnu"` //当日新增角色数
	MacDnu  int64  `json:"macDnu"`  //当日新增设备数
}

type DayDauData struct {
	DimDay  string `json:"dimDay"`
	Dau     int64  `json:"dau"`     //日活跃（登录过）玩家数
	RoleDau int64  `json:"roleDau"` //日活跃（登录过）角色数
	MacDau  int64  `json:"macDau"`  //日活跃（登录过）设备数
}

type DayRemainData struct {
	DimDay  string `json:"dimDay"`
	Num     int64
	RoleNum int64
	MacNum  int64
}

type DayRechargeData struct {
	DimDay        string  `json:"dimDay"`
	TotalIncome   float32 `json:"totalIncome"`   //当日总收入（货币）
	RechargeUsers int64   `json:"rechargeUsers"` //当日充值人数
}

type DayNewRechargeData struct {
	DimDay            string  `json:"dimDay"`
	NewRechargeIncome float32 `json:"newRechargeIncome"` //新增充值收入（货币）
	NewRechargeUsers  int64   `json:"newRechargeUsers"`  //新增充值用户
}

type PlatDnuData struct {
	Plat    string `json:"plat"`
	Dnu     int64  `json:"dnu"`     //平台新增账号数
	RoleDnu int64  `json:"roleDnu"` //平台新增角色数
	MacDnu  int64  `json:"macDnu"`  //平台新增设备数
}

type PlatRemainData struct {
	Plat    string `json:"plat"`
	Num     int64
	RoleNum int64
	MacNum  int64
}

type PlatDauData struct {
	Plat    string `json:"plat"`
	Dau     int64  `json:"dau"`     //日活跃（登录过）玩家数
	RoleDau int64  `json:"roleDau"` //日活跃（登录过）角色数
	MacDau  int64  `json:"macDau"`  //日活跃（登录过）设备数
}

type OnlineData struct {
	Time    int `json:"time"`    //时间点
	RoleNum int `json:"rolenum"` //数量
}

type CurOnlineData struct {
	ServerID   string `json:"serverid"`   //服务器id
	ServerName string `json:"servername"` //服务器名字
	RoleNum    int    `json:"rolenum"`    //在线数量
}

type AvgOnlineData struct {
	Plat      string `json:"plat"`      //服务器id
	Date      string `json:"date"`      //日期
	OnlineMin int    `json:"onlinemin"` //在线数量
}

type LeagueLineData struct {
	Time      int `json:"time"`      //时间点
	LeagueNum int `json:"leaguenum"` //数量
}
