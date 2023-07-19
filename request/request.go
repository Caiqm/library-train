package request

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Http struct {
}

// post请求
func (h Http) PostRequestWithByte(uri string, data []byte, contentType string) (body []byte, err error) {
	resp, err := http.Post(uri, contentType, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("请求失败")
		return
	}
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)
	return
}

// post请求
func (h Http) PostRequest(uri string, data url.Values, contentType string) (body []byte, err error) {
	resp, err := http.Post(uri, contentType, strings.NewReader(data.Encode()))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("请求失败")
		return
	}
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)
	return
}

// get请求
func (h Http) GetRequest(uri string) (body []byte, err error) {
	resp, err := http.Get(uri)
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		err = errors.New("请求失败")
		return
	}
	defer resp.Body.Close()
	body, _ = io.ReadAll(resp.Body)
	return
}
