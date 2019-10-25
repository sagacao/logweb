package ini

import (
	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/ini"
)

var Rechargeform struct{
	Info map[string]string
}

func initRechargeform() error {
	Rechargeform.Info = make(map[string]string)

	cfg, err := ini.Load("config/recharge.ini")
	section, err := cfg.GetSection("recharge")
	if section == nil || err != nil {
		return err
	}
	keys := section.KeyStrings()

	for _, k := range keys {
		v := section.Key(k).String()
		Rechargeform.Info[k] = v
	}

	faygo.Infof("initRechargeform : %v", Rechargeform)
	return nil
}