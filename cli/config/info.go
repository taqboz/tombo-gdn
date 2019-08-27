package config

import (
	"encoding/json"
	"io/ioutil"
)

type CheckItem struct {
	Tag string `json:"tag"`
	Target string `json:"info"`
	Attr map[string]string `json:"attr"`
	Length CheckLength `json:"length"`
	UseKws CheckLength `json:"use_kws"`
	NotPermit []string `json:"not_permit"`
	NotInclude []string `json:"not_include"`
	MultipleNum MultipleLength `json:"multiple_num"`
	MultipleDup MultipleLength  `json:"multiple_dup"`
}

type CheckLength struct {
	Check string `json:"check"`
	Min int `json:"min"`
	Max int `json:"max"`
}

type MultipleLength struct {
	Check string `json:"check"`
	Min int `json:"min"`
	Max int `json:"max"`
	Split string `json:"split"`
}

func CheckInit(file string) ([]CheckItem, error) {
	var l []CheckItem
	ignore, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(ignore, &l); err != nil {
		return nil, err
	}

	return l, nil
}
