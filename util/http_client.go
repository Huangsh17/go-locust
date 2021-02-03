package util

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type Method interface {
	Get(url string) (string, error)
	Post(url, body string) (string, error)
}

type HttpClient struct {
}

func (hc *HttpClient) Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		Sugar.Errorw("get requests fail ", "error", err)
	}
	resp, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return string(resp), err
}

func (hc *HttpClient) Post(url string, body string) (string, error) {
	contentType := "application/json"
	resp, err := http.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		Sugar.Errorw("post requests fail ", "error", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		Sugar.Errorw("io requests fail ", "error", err)
	}
	return string(b), err
}
