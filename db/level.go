package db

import (
	"fmt"
	"logweb/models/level"
	"time"

	"github.com/henrylee2cn/faygo"
)

func GetLevelDisData(zeusid string, stm, etm time.Time, prof int, results *[]*level.LevelDisData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	getnewlyleveldisdata(engine, stm, etm, prof, results)

	return nil

}

func GetLevelLoseData(zeusid string, stm, etm time.Time, prof int, results *[]*level.LevelLoseData) error {
	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	getnewlylevellosedata(engine, stm, etm, prof, results)
	return nil

}
