package main

import (
	"encoding/base64"
	"net/http"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func redirectPolicyFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
	return nil
}

func main() {
	client := &http.Client{
		Jar:           cookieJar,
		CheckRedirect: redirectPolicyFunc,
	}

	req, err := http.NewRequest("GET", "http://localhost/", nil)
	req.Header.Add("Authorization", "Basic "+basicAuth("username1", "password123"))
	resp, err := client.Do(req)
}
