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

func Get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
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
