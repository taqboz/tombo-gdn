package check

import (
	"log"
	"reflect"
	"testing"

	"github.com/taqboz/tombo/cli/config"
)

func TestDuplicateInpage(t *testing.T) {
	tests := []struct {
		in  []string
		out map[string]int
	}{
		{[]string{"aa", "bb", "aa", "bb", "cc"}, map[string]int{"aa": 2, "bb": 2}},
		{[]string{"ああ", "いい", "ああ", "ええ"}, map[string]int{"ああ": 2}},
		{[]string{"aa","ああ","a","bb"}, map[string]int{}},
		{[]string{}, nil},
	}

	for _, v := range tests {
		testTag := ErrTag{}
		err := duplicateInPage(&testTag, v.in)
		if err != nil {
			log.Fatal(err)
		}

		if !reflect.DeepEqual(v.out, testTag.DuplicateInPage) {
			t.Fatal("failed test")
		}
	}
}

func TestDuplicateInContent(t *testing.T) {
	m := config.MultipleContent{DuplicateInContent:true, Split:","}

	tests := []struct {
		in  string
		out map[string]int
	}{
		{"aa,bb,aa,bb,cc", map[string]int{"aa": 2, "bb": 2}},
		{"ああ,いい,ああ,ええ", map[string]int{"ああ": 2}},
		{"aa,ああ,a,bb", map[string]int{}},
		{"", nil},
	}

	for _, v := range tests {
		testTag := ErrTag{}
		err := duplicateInContent(&testTag, &m, v.in)
		if err != nil {
			log.Fatal(err)
		}

		if !reflect.DeepEqual(v.out, testTag.DuplicateInContent) {
			t.Fatal("failed test")
		}
	}

}

func TestDuplicate(t *testing.T) {
	c := []*config.Tag{
		{Tag: "title", Duplicate:true},
		{Tag: "meta", Target: "contents", Attr: map[string]string{"name": "description"}, Duplicate:true},
	}

	tests := []struct{
		cont string
		url string
	}{
		{"テストテスト", "http://abc.com/edf"},
		{"テストテスト", "http://abc.com/edf"},
		{"テストテスト", "http://abc.com/gh"},
		{"テスト", "http://abc.com/edf"},
		{"テスト", "http://abc.com/gh"},
	}

	testOut := []map[string]map[string][]string{
		{"title.|": {"テストテスト":{"http://abc.com/edf","http://abc.com/edf","http://abc.com/gh"},
			"テスト":{"http://abc.com/edf", "http://abc.com/gh"}}},
		{"meta.contents|name:description,": {"テストテスト":{"http://abc.com/edf","http://abc.com/edf","http://abc.com/gh"},
			"テスト":{"http://abc.com/edf", "http://abc.com/gh"}}},
	}


	for k, v := range c {
		for _, v2 := range tests {
			duplicate(v, v2.cont, v2.url)
		}
		if !reflect.DeepEqual(ErrDupulicateList, testOut[k]) {
			t.Fatal("failed test")
		}
		ErrDupulicateList = map[string]map[string][]string{}
	}
}
