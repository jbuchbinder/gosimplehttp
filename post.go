package gosimplehttp

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

// SimplePostJson is a wrapper for a ridiculously simple POST
// request for JSON content.
func SimplePostJson(url, name, filename string) (code int, resp []byte, head http.Header, err error) {
	return SimplePostJsonWithAuth(url, name, filename, "", "")
}

// SimplePostJsonNoFile is a wrapper for a ridiculously simple POST
// request for JSON content, pulled from local content
func SimplePostNoFileJson(url, name string, content []byte) (code int, resp []byte, head http.Header, err error) {
	return SimplePostJsonNoFileWithAuth(url, name, content, "", "")
}

// SimplePostJsonNoFileWithAuth is a wrapper for a ridiculously simple
// POST request for JSON content with BASIC authentication.
func SimplePostJsonNoFileWithAuth(url, name string, content []byte, username, password string) (code int, resp []byte, head http.Header, err error) {
	c := NewClient()
	if username != "" && password != "" {
		c.SetAuthentication(username, password)
	}
	fComp := PostValue(name, string(content))
	comp := []MultipartComponenter{fComp}
	return c.DoPost(url, comp, nil)
}

// SimplePostJsonWithAuth is a wrapper for a ridiculously simple
// POST request for JSON content with BASIC authentication.
func SimplePostJsonWithAuth(url, name, filename, username, password string) (code int, resp []byte, head http.Header, err error) {
	c := NewClient()
	if username != "" && password != "" {
		c.SetAuthentication(username, password)
	}
	fComp := PostFile(name, filename, "application/json")
	comp := []MultipartComponenter{fComp}
	return c.DoPost(url, comp, nil)
}

// SimplePostXml is a wrapper for a ridiculously simple POST
// request for XML content.
func SimplePostXml(url, name, filename string) (code int, resp []byte, head http.Header, err error) {
	return SimplePostXmlWithAuth(url, name, filename, "", "")
}

// SimplePostXmlWithAuth is a wrapper for a ridiculously simple
// POST request for XML content with BASIC authentication.
func SimplePostXmlWithAuth(url, name, filename, username, password string) (code int, resp []byte, head http.Header, err error) {
	c := NewClient()
	if username != "" && password != "" {
		c.SetAuthentication(username, password)
	}
	fComp := PostFile(name, filename, "application/xml")
	comp := []MultipartComponenter{fComp}
	return c.DoPost(url, comp, nil)
}

// DoPost executes a POST request with the specified criteria.
func (s *SimpleHttpClient) DoPost(url string, components []MultipartComponenter, headers map[string]string) (code int, resp []byte, head http.Header, err error) {
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
	resp, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	code = res.StatusCode

	return
}
