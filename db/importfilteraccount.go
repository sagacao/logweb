package db

import (
	"fmt"

	"github.com/henrylee2cn/faygo"
)

func ImportFilterAccountIDs(zeusid string, accids []string) error {

	engine := reader.getReader(zeusid)
	if engine == nil {
		faygo.Warningf("Create db engine '%s' ", zeusid)
		return fmt.Errorf("大区不存在: %v", zeusid)
	}

	if len(accids) <= 0 {
		return fmt.Errorf("账号ID太少了")
	}
	sql := fmt.Sprintf("insert into log_filter values")

	for k, v := range accids {

		if k > 0 {
			sql += ","
		}
		sql = sql + "(" + v + ")"
	}
	sql += ";"

	_, err := engine.Query(sql)

	return err
}
