package http

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDoRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			return
		},
	))
	defer ts.Close()

	req, err := http.NewRequest("GET", ts.URL, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	res, err := DoRequest(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if res.StatusCode != 200 {
		fmt.Println("test is failed")
		os.Exit(1)
	}
}
