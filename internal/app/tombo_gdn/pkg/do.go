package pkg

import "net/http"

// リクエストの実行
func DoRequest(req *http.Request) (*http.Response, error){
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
