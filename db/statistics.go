package db

import (
	"fmt"
	"logweb/models/base"
	"logweb/models/statistics"
	"logweb/utils"
	"time"

	"github.com/henrylee2cn/faygo"
)

func mergeRemainData(day int, v *statistics.RemainData, accountval, roleval, macval string) {
	switch day {
	case 1:
		v.AccountNdrr1 = accountval
		v.RoleNdrr1 = roleval
		v.MacNdrr1 = macval
	case 2:
		v.AccountNdrr3 = accountval
		v.RoleNdrr3 = roleval
		v.MacNdrr3 = macval
	case 3:
		v.AccountNdrr4 = accountval
		v.RoleNdrr4 = roleval
		v.MacNdrr4 = macval
	case 4:
		v.AccountNdrr5 = accountval
		v.RoleNdrr5 = roleval
		v.MacNdrr5 = macval
	case 5:
		v.AccountNdrr6 = accountval
		v.RoleNdrr6 = roleval
		v.MacNdrr6 = macval
	case 6:
		v.AccountNdrr7 = accountval
		v.RoleNdrr7 = roleval
		v.MacNdrr7 = macval
	}
}

func GetRemainData(zeusid string, stm, etm time.Time, results *[]*statistics.RemainData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	dnu := make(map[string]*base.DayDnuData)
	getnewlydata(engine, stm, etm, dnu)
	for _, v := range *results {
		key := v.DimDay
		dnudata, ok := dnu[key]
		if ok {
			v.AccountDnu = dnudata.Dnu
			v.RoleDnu = dnudata.RoleDnu
			v.MacDnu = dnudata.MacDnu
		}
	}

	searchtm := stm
	stm = stm.Add(time.Hour * 24)
	daycounter := 0
	for stm.Before(etm) {

		daycounter++
		realtm := searchtm
		dayremain := make(map[string]*base.DayRemainData)
		for realtm.Before(etm) {
			refertm := realtm.Add(time.Hour * 24 * time.Duration(daycounter))
			getdailyremain(engine, refertm, realtm, dayremain)
			realtm = realtm.Add(time.Hour * 24)
		}

		//faygo.Warningf("=====> day:[%v] time:%v||%v ", daycounter, stm, etm)
		for _, v := range *results {
			key := v.DimDay
			remaindata, ok := dayremain[key]
			if ok {
				accval := utils.FormatPercent(float32(remaindata.Num), float32(v.AccountDnu))
				roleval := utils.FormatPercent(float32(remaindata.RoleNum), float32(v.RoleDnu))
				macval := utils.FormatPercent(float32(remaindata.MacNum), float32(v.MacDnu))
				mergeRemainData(daycounter, v, accval, roleval, macval)
			}
		}

		stm = stm.Add(time.Hour * 24)

		if daycounter > 7 {
			break
		}
	}

	return nil
}

func GetPlatRemainData(zeusid string, stm, etm time.Time, platremainretdata map[string]*statistics.RemainData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	//platremainretdata := make(map[string]*statistics.RemainData)

	platdnu := make(map[string]*base.PlatDnuData)
	getnewlyplatdata(engine, stm, platdnu)
	for k, v := range platdnu {
		platdnudata, ok := platremainretdata[k]
		if ok {
			platdnudata.AccountDnu = v.Dnu
			platdnudata.RoleDnu = v.RoleDnu
			platdnudata.MacDnu = v.MacDnu
		} else {
			platremainretdata[k] = &statistics.RemainData{
				DimDay:     k,
				AccountDnu: v.Dnu,
				RoleDnu:    v.RoleDnu,
				MacDnu:     v.MacDnu,
			}
		}
	}

	refertm := stm
	stm = stm.Add(time.Hour * 24)
	daycounter := 0
	for stm.Before(etm) {
		daycounter++
		platremain := make(map[string]*base.PlatRemainData)
		err := getplatremain(engine, stm, refertm, platremain)
		if err == nil {
			for k, _ := range platdnu {
				platdnudata, ok1 := platremainretdata[k]
				platremaindata, ok2 := platremain[k]
				if ok1 && ok2 {
					//faygo.Debugf("===========>  %v %v", platremaindata.Num, platdnudata.AccountDnu)
					accval := utils.FormatPercent(float32(platremaindata.Num), float32(platdnudata.AccountDnu))
					roleval := utils.FormatPercent(float32(platremaindata.RoleNum), float32(platdnudata.RoleDnu))
					macval := utils.FormatPercent(float32(platremaindata.MacNum), float32(platdnudata.MacDnu))
					mergeRemainData(daycounter, platdnudata, accval, roleval, macval)
				}
			}
		}
		stm = stm.Add(time.Hour * 24)
	}

	return nil
}

func GetDauData(zeusid string, stm, etm time.Time, results *[]*base.DayDauData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	dau := make(map[string]*base.DayDauData)
	for stm.Before(etm) {
		getdailydau(engine, stm, dau)
		stm = stm.Add(time.Hour * 24)
	}

	for k, v := range dau {
		*results = append(*results, &base.DayDauData{
			DimDay:  k,
			Dau:     v.Dau,
			RoleDau: v.RoleDau,
			MacDau:  v.MacDau,
		})
	}

	return nil
}

func GetDauPlatData(zeusid string, stm, etm time.Time, results *[]*base.PlatDauData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	for stm.Before(etm) {
		dau := make(map[string]*base.PlatDauData)
		getplatdau(engine, stm, dau)
		for k, v := range dau {
			*results = append(*results, &base.PlatDauData{
				Plat:    k,
				Dau:     v.Dau,
				RoleDau: v.RoleDau,
				MacDau:  v.MacDau,
			})
		}

		stm = stm.Add(time.Hour * 24)
	}

	return nil
}

func GetRunoffData(zeusid string, stm, etm time.Time, querytype string, daycounter int, results *[]*statistics.RunOffData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	getrunoff(engine, stm, querytype, daycounter, results)

	return nil
}
