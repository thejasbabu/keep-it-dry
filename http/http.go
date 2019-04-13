package http

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"

	er "github.com/thejasbabu/keep-it-dry/error"
)

// Get request for a give path with headers
func Get(path string, headers map[string]string) ([]byte, error) {
	client := &http.Client{}
	return getResponse(client, path, headers)
}

// InsecureGet request bypassing the certificate check for a given path with headers
func InsecureGet(path string, headers map[string]string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return getResponse(client, path, headers)
}

func getResponse(client *http.Client, path string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", path, nil)
	if er.IsError(err) {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if er.IsError(err) {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if er.IsError(err) {
		return nil, err
	}
	return body, nil
}
