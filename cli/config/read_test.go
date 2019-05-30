package config

import "testing"

func TestReadJson(t *testing.T) {
	tag, cp, cl, clt := ReadJson("../../testdata/config.json")
	if tag == nil || cp == 10 || cl == 10 || clt == 1000 {
		t.Fatal("failed test")
	}
}