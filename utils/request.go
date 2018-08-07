package utils

import (
	"bytes"
	"encoding/json"
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

	req.Header.Set("Accept", "application/json")

	for _, v := range externalHeader {
		for k := range v {
			req.Header.Set(k, v[k])
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var bf bytes.Buffer
	io.Copy(&bf, resp.Body)
	return bf.Bytes(), nil
}

func Post(url string, body interface{}, externalHeader ...map[string]string) ([]byte, error) {
	data, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")

	for _, v := range externalHeader {
		for k := range v {
			req.Header.Set(k, v[k])
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var bf bytes.Buffer
	io.Copy(&bf, resp.Body)
	return bf.Bytes(), nil
}
