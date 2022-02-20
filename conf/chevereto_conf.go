package conf

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

type ApiConf struct {
	Key string
	Url string
}

func CheveretoConf() ApiConf {
	cfg, err := ini.Load("conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	var apiConf ApiConf
	apiConf.Url = cfg.Section("chevereto").Key("url").String()
	apiConf.Key = cfg.Section("chevereto").Key("key").String()
	return apiConf
}
