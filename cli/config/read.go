package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"
)

// チェック設定用jsonファイルの読み込み
func ReadJson(file string) ([]*Tag, int, time.Duration) {
	// 設定ファイルの読み込み
	config_json, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	// jsonファイルのエンコーディング
	var config *Config
	if err := json.Unmarshal(config_json, &config); err != nil {
		log.Fatal(err)
	}
	return config.Tag, config.CheckPageParallel, config.GetLinksTimeSleep
}

func AsignConfig(file string) {
	Tags, CheckPageParallel, GetLinksTimeSleep = ReadJson(file)
}
