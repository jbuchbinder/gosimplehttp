package gosimplehttp

import (
	"net/http"
)

const (
	REQUEST_DELETE  = "DELETE"
	REQUEST_GET     = "GET"
	REQUEST_OPTIONS = "OPTIONS"
	REQUEST_PUT     = "PUT"
	REQUEST_POST    = "POST"
)

// SimpleHttpClient is the base type of the gosimplehttp client. It should
// ideally be instantiated by NewClient(), and should be called with the
// Set* methods, then have calls executed with Do* methods.
type SimpleHttpClient struct {
	client      *http.Client
	username    string
	password    string
	initialized bool
	cookies     []*http.Cookie
}

// NewClient instantiates a pointer to a SimpleHttpClient object, which is
// the base of all client operations.
func NewClient() *SimpleHttpClient {
	c := &SimpleHttpClient{}
	c.cookies = []*http.Cookie{}
	return c
}

func (s *SimpleHttpClient) init() {
	tr := &http.Transport{
		//TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	s.client = &http.Client{Transport: tr}
}

// AddCookie adds a pointer to an http.Cookie object to the list
// of cookies being passed to subsequent client requests.
func (s *SimpleHttpClient) AddCookie(cookie *http.Cookie) {
	s.cookies = append(s.cookies, cookie)
}

// AddCookie adds list of pointers to http.Cookie objects to the
// list of cookies being passed to subsequent client requests.
func (s *SimpleHttpClient) AddCookies(cookielist []*http.Cookie) {
	for _, v := range cookielist {
		s.cookies = append(s.cookies, v)
	}
}

// GetClient retrieves a pointer to the underlying http.Client instance.
func (s *SimpleHttpClient) GetClient() *http.Client {
	return s.client
}

// SetAuthentication sets BASIC authentication username and password for
// all further client requests.
func (s *SimpleHttpClient) SetAuthentication(u, p string) {
	s.username = u
	s.password = p
}
