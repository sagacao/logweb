package statistics

import "time"

type RemainData struct {
	DimDay       string `json:"dimDay"`
	AccountDnu   int64  `json:"accountDnu"`   //当日新增账号数
	RoleDnu      int64  `json:"roleDnu"`      //当日新增角色数
	MacDnu       int64  `json:"macDnu"`       //当日新增设备数
	AccountNdrr1 string `json:"accountNdrr1"` //次日留存率(百分数)
	RoleNdrr1    string `json:"roleNdrr1"`    //次日留存率(百分数)
	MacNdrr1     string `json:"macNdrr1"`     //次日留存率(百分数)
	AccountNdrr3 string `json:"accountNdrr3"` //3日留存率(百分数)
	RoleNdrr3    string `json:"roleNdrr3"`    //3日留存率(百分数)
	MacNdrr3     string `json:"macNdrr3"`     //3日留存率(百分数)
	AccountNdrr4 string `json:"accountNdrr4"` //4日留存率(百分数)
	RoleNdrr4    string `json:"roleNdrr4"`    //4日留存率(百分数)
	MacNdrr4     string `json:"macNdrr4"`     //4日留存率(百分数)
	AccountNdrr5 string `json:"accountNdrr5"` //5日留存率(百分数)
	RoleNdrr5    string `json:"roleNdrr5"`    //5日留存率(百分数)
	MacNdrr5     string `json:"macNdrr5"`     //5日留存率(百分数)
	AccountNdrr6 string `json:"accountNdrr6"` //6日留存率(百分数)
	RoleNdrr6    string `json:"roleNdrr6"`    //6日留存率(百分数)
	MacNdrr6     string `json:"macNdrr6"`     //6日留存率(百分数)
	AccountNdrr7 string `json:"accountNdrr7"` //7日留存率(百分数)
	RoleNdrr7    string `json:"roleNdrr7"`    //7日留存率(百分数)
	MacNdrr7     string `json:"macNdrr7"`     //7日留存率(百分数)
}

func InitRemainData(stm, etm time.Time, result *[]*RemainData) {
	tmcursor := stm
	for tmcursor.Before(etm) {
		*result = append(*result, &RemainData{
			DimDay:       tmcursor.Format("2006-01-02"),
			AccountDnu:   0,
			RoleDnu:      0,
			MacDnu:       0,
			AccountNdrr1: "0",
			RoleNdrr1:    "0",
			MacNdrr1:     "0",
			AccountNdrr3: "0",
			RoleNdrr3:    "0",
			MacNdrr3:     "0",
			AccountNdrr4: "0",
			RoleNdrr4:    "0",
			MacNdrr4:     "0",
			AccountNdrr5: "0",
			RoleNdrr5:    "0",
			MacNdrr5:     "0",
			AccountNdrr6: "0",
			RoleNdrr6:    "0",
			MacNdrr6:     "0",
			AccountNdrr7: "0",
			RoleNdrr7:    "0",
			MacNdrr7:     "0",
		})
		tmcursor = tmcursor.Add(time.Hour * 24)
	}
}

type RunOffData struct {
	Plat    int    `json:"plat"`
	Number  string `json:"number"`
	RoleNum int64  `json:"rolenum"`
}
