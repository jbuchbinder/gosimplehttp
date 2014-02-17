package gosimplehttp

import (
	"io/ioutil"
	"net/http"
)

func (s *simpleHttpClient) DoPost(url string, params map[string][]byte, headers map[string]string) (code int, resp []byte, err error) {
	if !s.initialized {
		s.init()
	}
	req, err := http.NewRequest(REQUEST_POST, url, nil)
	if err != nil {
		return
	}
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
	if s.username != "" && s.password != "" {
		req.SetBasicAuth(s.username, s.password)
	}
	res, err := s.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	code = res.StatusCode
	resp = body

	return
}
