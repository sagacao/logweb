package ini

import (
	"github.com/henrylee2cn/faygo"
	"github.com/henrylee2cn/ini"
)

var Platform struct {
	Info map[string]string
}

func initPlatform() error {
	Platform.Info = make(map[string]string)
	
	cfg, err := ini.Load("config/platform.ini")
	section, err := cfg.GetSection("platform")
	if section == nil || err != nil {
		return err
	}
	keys := section.KeyStrings()

	for _, k := range keys {
		v := section.Key(k).String()
		Platform.Info[k] = v
	}

	faygo.Infof("InitPlatform : %v", Platform)
	return nil
}