package config

import (
	"encoding/json"
	"github.com/taqboz/tombo_gdn/internal/app/tombo_gdn/pkg"
	"io/ioutil"
)

type CheckItem struct {
	Tag string `json:"tag"`
	Target string `json:"target"`
	Attr map[string]string `json:"attr"`
	Length CheckLength `json:"length"`
	UseKws CheckLength `json:"use_kws"`
	NotPermit []string `json:"not_permit"`
	NotInclude []string `json:"not_include"`
}

type CheckLength struct {
	Check string `json:"check"`
	Min int `json:"min"`
	Max int `json:"max"`
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

	conf := ignoreTags(l)

	return conf, nil
}


func ignoreTags(bf []CheckItem) []CheckItem {
	af := []CheckItem{}
	for _, v := range bf {
		 if !pkg.IsContain(Ignore.IgnoreTags, v.Tag) {
		 	af = append(af, v)
		 }
	}

	return af
}
