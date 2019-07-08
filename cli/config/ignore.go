package config

import "github.com/go-ini/ini"

var Ignore = []string{}

func IgnoreInit() error {
	cfg, err := ini.Load("config/auth.ini")
	if err != nil {
		return err
	}
	cfg.Section("ignore").Key("paths").String()

	return nil
}
