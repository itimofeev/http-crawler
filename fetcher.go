package main

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type httpFetcher struct {
	client http.Client
}

func newHTTPFetcher() *httpFetcher {
	return &httpFetcher{
		client: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (f *httpFetcher) Fetch(url string) urlAndHash {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	getResp, err := f.client.Get(url)
	if err != nil {
		return urlAndHash{
			url:  url,
			hash: "error making http request: " + err.Error(),
		}
	}
	defer func() {
		_ = getResp.Body.Close()
	}()
	bodyBytes, err := ioutil.ReadAll(getResp.Body)
	if err != nil {
		return urlAndHash{
			url:  "url",
			hash: "error reading response body",
		}
	}
	sum := md5.Sum(bodyBytes)
	return urlAndHash{
		url:  url,
		hash: hex.EncodeToString(sum[:]),
	}
}
