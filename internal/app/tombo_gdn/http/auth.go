package http

import (
	"encoding/base64"
	"net/http"
)

func BasicAuth(req *http.Request, username string, password string) *http.Request {
	req.Header.Add("Authorization","Basic " + basicAuth(username,password))
	return req
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
