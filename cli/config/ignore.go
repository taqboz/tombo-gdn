package config

import (
	"encoding/json"
	"io/ioutil"
)

var Ignore IgnoreConfig

type IgnoreConfig struct {
	IgnorePaths []string `json:"ignore_paths"`
	IgnoreTags []string `json:"ignore_tags"`
}

func IgnoreInit(file string) error {
	ignore, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(ignore, &Ignore); err != nil {
		return err
	}

	return nil
}
