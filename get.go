package gosimplehttp

import (
	"io/ioutil"
	"net/http"
)

// SimpleGet is a wrapper for a ridiculously simple GET request.
func SimpleGet(url string) (code int, resp []byte, head http.Header, err error) {
	return SimpleGetWithAuth(url, "", "")
}

// SimpleGetTimeout is a wrapper for a ridiculously simple GET request.
func SimpleGetTimeout(url string, timeout int) (code int, resp []byte, head http.Header, err error) {
	return SimpleGetWithAuthTimeout(url, "", "", timeout)
}

// SimpleGetWithAuth is a wrapper for a ridiculously simple GET request
// with BASIC authentication.
func SimpleGetWithAuth(url, username, password string) (code int, resp []byte, head http.Header, err error) {
	c := NewClient()
	c.Timeout = 30
	if username != "" && password != "" {
		c.SetAuthentication(username, password)
	}
	return c.DoGet(url, nil)
}

// SimpleGetWithAuthTimeout is a wrapper for a ridiculously simple GET request
// with BASIC authentication and a specified timeout.
func SimpleGetWithAuthTimeout(url, username, password string, timeout int) (code int, resp []byte, head http.Header, err error) {
	c := NewClient()
	c.Timeout = timeout
	if username != "" && password != "" {
		c.SetAuthentication(username, password)
	}
	return c.DoGet(url, nil)
}

// DoGet executes a GET with the specified criteria.
func (s *SimpleHttpClient) DoGet(url string, headers map[string]string) (code int, resp []byte, head http.Header, err error) {
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
	head = res.Header
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	code = res.StatusCode
	resp = body

	return
}
