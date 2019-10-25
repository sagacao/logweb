package handler

import (
	"encoding/csv"
	"io"
	"log"
	"logweb/db"

	"github.com/henrylee2cn/faygo"
)

/*
ImportFilterAccount
*/
var ImportFilterAccount = faygo.HandlerFunc(func(ctx *faygo.Context) error {

	zeusid := ctx.Param("zeusid")
	formFile, _, err := ctx.FormFile("xlsfile")
	if err != nil {
		log.Printf("Get form file failed: %s\n", err)
		return nil
	}

	defer formFile.Close()

	accountids := []string{}

	reader := csv.NewReader(formFile)
	record := []string{}
	for {
		record, err = reader.Read()
		if err == io.EOF {
			break
		} else if nil != err {
			return nil
		}

		accountids = append(accountids, record[0])
	}

	db.ImportFilterAccountIDs(zeusid, accountids)

	return ctx.JSONMsg(200, 200, "成功")
})
