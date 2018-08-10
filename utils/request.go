package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/zcong1993/badge-service/cache"
	"github.com/zcong1993/utils/iocopy"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: time.Second * 6,
}

// ErrMaxParallel is error msg of request hit MaxParallel
var ErrMaxParallel = errors.New("request hit MaxParallel")

var pooledCopy = iocopy.DefaultPolledIoCopyFunc

// Get is helper function for making get request
func Get(url string, externalHeader ...map[string]string) ([]byte, error) {
	// check if request hit MaxParallel
	if cache.IsBurst(url) {
		return nil, ErrMaxParallel
	}
	defer cache.Release(url)

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
	pooledCopy(&bf, resp.Body)
	return bf.Bytes(), nil
}

// Post is helper function for making post request
func Post(url string, body interface{}, externalHeader ...map[string]string) ([]byte, error) {
	// check if request hit MaxParallel
	if cache.IsBurst(url) {
		return nil, ErrMaxParallel
	}
	defer cache.Release(url)

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
	pooledCopy(&bf, resp.Body)
	return bf.Bytes(), nil
}
