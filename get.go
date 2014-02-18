package gosimplehttp

import (
	"io/ioutil"
	"net/http"
)

// DoGet executes a GET with the specified criteria.
func (s *SimpleHttpClient) DoGet(url string, headers map[string]string) (code int, resp []byte, err error) {
	if !s.initialized {
		s.init()
	}
	req, err := http.NewRequest(REQUEST_GET, url, nil)
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
	if s.cookies != nil && len(s.cookies) > 0 {
		for _, v := range s.cookies {
			req.AddCookie(v)
		}
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
