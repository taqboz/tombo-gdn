package target

import (
	"log"
	"testing"
)

// ローカルホストサーバーを起動してsitemap.xmlを表示
func TestInputTargetInfo(t *testing.T) {
	t.Helper()
	test := []struct{
		url string
		host string
	} {
		{"http://localhost:3000/sitemap.xml","http://localhost:3000/"},
	}

	for _, v := range test {
		if err := InputTargetInfo(v.url); err != nil {
			log.Fatal(err)
		} else {
			if Input != v.url {
				t.Fatal("failed test | InputUrl:",v.url, Input)
			}
			if Host != v.host {
				t.Fatal("failed test | HostUrl:",v.host, Host)
			}
		}
	}
}