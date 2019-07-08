package config

import "github.com/go-ini/ini"

type BasicAuthList struct {
	UserName  string
	Passwords string
}

var BasicAuth  BasicAuthList

func AuthInit() error {
	cfg, err := ini.Load("config/auth.ini")
	if err != nil {
		return err
	}

	BasicAuth =  BasicAuthList{
		UserName: cfg.Section("basic_auth").Key("username").String(),
		Passwords: cfg.Section("basic_auth").Key("passwords").String(),
	}

	return nil
}
