package gosimplehttp

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func (s *simpleHttpClient) DoPost(url string, components []multipartComponenter, headers map[string]string) (code int, resp []byte, err error) {
	if !s.initialized {
		s.init()
	}

	// Prep content
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for c := range components {
		components[c].SetWriter(*writer)
		components[c].Encode()
	}
	err = writer.Close()
	if err != nil {
		return
	}

	req, err := http.NewRequest(REQUEST_POST, url, body)
	if err != nil {
		return
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	req.Header.Set("Content-type", writer.FormDataContentType())
	if s.username != "" && s.password != "" {
		req.SetBasicAuth(s.username, s.password)
	}
	res, err := s.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	resp, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	code = res.StatusCode

	return
}
