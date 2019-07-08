package pkg

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http/httptest"
	"os"
	"testing"
)

func TestBasicAuth(t *testing.T) {
	tests := []struct {
		user string
		pass string
	}{
		{"admin","password"},
		{"admn","password"},
		{"01234","password"},
		{"01234","pass0000"},
		{"あいう", "かきくけ"},
	}

	for k, v := range tests {
		req := httptest.NewRequest("GET", "http://example.com/", nil)
		req = BasicAuth(req, v.user, v.pass)
		user, pass, ok := req.BasicAuth()

		// Basic認証のヘッダ解析に失敗した場合
		if !ok {
			err := errors.New("Parsing basic authorication is failed")
			fmt.Println(err)
			os.Exit(1)
		}

		//
		if v.user != user || v.pass != pass {
			err := fmt.Errorf("username: %s, password: %s|test %d", v.user, v.pass, k)
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func TestBasicAuth2(t *testing.T) {
	tests := []struct {
		user string
		pass string
	}{
		{"admin","password"},
		{"admn","password"},
		{"01234","password"},
		{"01234","pass0000"},
		{"あいう", "かきくけ"},
	}

	for k, v := range tests {
		st := v.user+":"+v.pass
		en := basicAuth(v.user, v.pass)
		data, err := base64.StdEncoding.DecodeString(en)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if string(data) != st {
			err := fmt.Errorf("username: %s, password: %s|test %d", v.user, v.pass, k)
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
