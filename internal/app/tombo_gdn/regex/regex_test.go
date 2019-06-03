package regex

import (
	"strconv"
	"testing"
)

func TestHostURL(t *testing.T) {
	regexTests := []struct{
		// 主にFindAllStringで使用する正規表現なので以下の構成
		in string
		out string
	}{
		// http(s)のURLのみ抽出できるか確認
		{"http://abc.com", "http://abc.com"},
		{"http://abc.com/", "http://abc.com/"},
		{"http://www.abc.com/", "http://www.abc.com/"},
		{"https://abc.com/", "https://abc.com/"},
		{"https://abc.co.jp/", "https://abc.co.jp/"},
		{"http://abc.com/def", "http://abc.com/"},
		{"http://abc.com/def/", "http://abc.com/"},

		{"", ""},
		{"abc", ""},
		{"www.abc.com/", ""},
		{"abc.com/", ""},
		{"abc@example.com", ""},
	}

	for i, test := range regexTests {
		s := HostURL.FindAllString(test.in, -1)
		if len(s) == 0 {
			if test.out != "" {
				t.Fatal("failed test", strconv.Itoa(i))
			}
		} else if s[0] != test.out {
			t.Fatal("failed test", strconv.Itoa(i))
		}
	}
}

func TestPageLink(t *testing.T) {
	regexTests := []struct{
		in string
		out bool
	}{
		// http(s)かwww.のURLを抽出できるか確認
		{"http://abc.com", true},
		{"http://abc.com/", true},
		{"http://www.abc.com/", true},
		{"https://abc.com/", true},
		{"https://abc.co.jp/", true},
		{"http://abc.com/def", true},
		{"http://abc.com/def/", true},

		{"abc.com", false},
		{"", false},
		{"abc", false},
		{"abc.com/", false},
		{"www.abc.com", false},
		{"www.abc.com/", false},
		{"tel:000-1234-5678", false},
		{"mailto:abc@example.com", false},
	}

	for i, test := range regexTests {
		if s := PageLink.MatchString(test.in); s != test.out {
			t.Fatal("failed test", strconv.Itoa(i))
		}
	}
}

func TestXML(t *testing.T)  {
	regexTests := []struct{
		in string
		out bool
	}{
		// http(s)かwww.のURLを抽出できるか確認
		{"http://abc.com/def.xml", true},
		{"http://www.abc.com/def.xml", true},
		{"https://abc.com/def.xml", true},
		{"https://abc.co.jp/def.xml", true},
		{"http://abc.com/def/ghi.xml", true},

		{"", false},
		{"abc", false},
		{"abc.com", false},
		{"abc.com/def.xml", false},
		{"www.abc.com/def.xml", false},
		{"tel:000-1234-5678", false},
		{"mailto:abc@example.com", false},
	}

	for i, test := range regexTests {
		if s := XML.MatchString(test.in); s != test.out {
			t.Fatal("failed test", strconv.Itoa(i))
		}
	}
}

func TestRelative(t *testing.T) {
	regexTests := []struct{
		in string
		out bool
	}{
		// http(s)かwww.のURLを抽出できるか確認
		{"/abc", true},
		{"/abc/def", true},
		{"/abc/#def", true},

		{"", false},
		{"abc", false},
		{"abc/def", false},

		{"abc.com", false},
		{"www.abc.com/", false},
		{"https://abc.com/", false},
		{"https://abc.co.jp/", false},
		{"http://abc.com/def", false},
		{"http://abc.com/def/", false},
		{"tel:000-1234-5678", false},
		{"mailto:abc@example.com", false},
	}

	for i, test := range regexTests {
		if s := Relative.MatchString(test.in); s != test.out {
			t.Fatal("failed test", strconv.Itoa(i))
		}
	}
}