package ini

import (
	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/ini"
)

var Itemchange struct {
	Info map[string]string
}

func initItemChange() error {

	Itemchange.Info = make(map[string]string)

	cfg, err := ini.Load("config/itemchange.ini")
	
	section, err := cfg.GetSection("itemchangereason")

	if section == nil || err != nil {
		return err
	}
	keys := section.KeyStrings()

	for _, k := range keys {

		v := section.Key(k).String()

		Itemchange.Info[k] = v

	}

	faygo.Infof("InitItemChange : %v", Itemchange)
	return nil
}