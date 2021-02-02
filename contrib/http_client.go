package contrib

import (
	"io"
	"net/http"
)

type Method interface {
	Get(url string) (resp *http.Response, err error)
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type HttpClient struct {
}

func (hc *HttpClient) Get(url string) (resp *http.Response, err error) {
	_, _ = http.Get(url)
	return nil, err
}

func (hc *HttpClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	_, _ = http.Post(url, contentType, body)
	return nil, err
}

func SetTask() {

}
