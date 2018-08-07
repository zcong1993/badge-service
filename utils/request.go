package utils

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 6,
}

// Get is helper function for making get request
func Get(url string, externalHeader ...map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for _, v := range externalHeader {
		for k := range v {
			req.Header.Set(k, v[k])
		}
	}

	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var bf bytes.Buffer
	io.Copy(&bf, resp.Body)
	return bf.Bytes(), nil
}
