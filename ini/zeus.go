package ini

import (
	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/ini"
)

var ZeusType struct {
	Info    map[string]string
	Ios     map[string]string
	Android map[string]string
}

func initZeusType() error {
	ZeusType.Info = make(map[string]string)
	ZeusType.Ios = make(map[string]string)
	ZeusType.Android = make(map[string]string)

	cfg, err := ini.Load("config/zeus.ini")
	section, err := cfg.GetSection("zeus")
	if section == nil || err != nil {
		return err
	}
	keys := section.KeyStrings()

	for _, k := range keys {
		v := section.Key(k).String()
		ZeusType.Info[k] = v
	}

	section, err = cfg.GetSection("ios")
	if section == nil || err != nil {
		return err
	}
	keys = section.KeyStrings()
	for _, k := range keys {
		v := section.Key(k).String()
		ZeusType.Ios[k] = v
	}

	section, err = cfg.GetSection("android")
	if section == nil || err != nil {
		return err
	}
	keys = section.KeyStrings()
	for _, k := range keys {
		v := section.Key(k).String()
		ZeusType.Android[k] = v
	}
	
	faygo.Infof("InitZeusType : %v", ZeusType)
	return nil
}
