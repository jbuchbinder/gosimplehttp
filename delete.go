package gosimplehttp

import (
	"io/ioutil"
	"net/http"
)

// SimpleDelete is a wrapper for a ridiculously simple DELETE
// request.
func SimpleDelete(url string) (code int, resp []byte, head http.Header, err error) {
	return SimpleDeleteWithAuth(url, "", "")
}

// SimpleDeleteWithAuth is a wrapper for a ridiculously simple
// DELETE request with BASIC authentication.
func SimpleDeleteWithAuth(url, username, password string) (code int, resp []byte, head http.Header, err error) {
	c := NewClient()
	if username != "" && password != "" {
		c.SetAuthentication(username, password)
	}
	return c.DoDelete(url, nil)
}

// DoDelete executes a DELETE with the specified criteria.
func (s *SimpleHttpClient) DoDelete(url string, headers map[string]string) (code int, resp []byte, head http.Header, err error) {
	if !s.initialized {
		s.init()
	}
	req, err := http.NewRequest(REQUEST_DELETE, url, nil)
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
	head = res.Header
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	code = res.StatusCode
	resp = body

	return
}
