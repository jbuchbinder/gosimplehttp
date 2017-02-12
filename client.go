package gosimplehttp

import (
	"crypto/tls"
	"net/http"
	"time"
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
	Timeout         int
	client          *http.Client
	username        string
	password        string
	initialized     bool
	cookies         []*http.Cookie
	tlsClientConfig *tls.Config
}

// NewClient instantiates a pointer to a SimpleHttpClient object, which is
// the base of all client operations.
func NewClient() *SimpleHttpClient {
	c := &SimpleHttpClient{}
	c.cookies = []*http.Cookie{}
	return c
}

// NewClientWithTlsParams instantiates a pointer to a SimpleHttpClient
// object, which is the base of all client operations, with specified
// TLS configuration.
func NewClientWithTlsParams(tlsConfig *tls.Config) *SimpleHttpClient {
	c := &SimpleHttpClient{}
	c.cookies = []*http.Cookie{}
	c.tlsClientConfig = tlsConfig
	return c
}

func (s *SimpleHttpClient) init() {
	tr := &http.Transport{
		DisableCompression: true,
	}
	if s.tlsClientConfig != nil {
		tr.TLSClientConfig = s.tlsClientConfig
	}
	s.client = &http.Client{Transport: tr, Timeout: time.Duration(s.Timeout) * time.Second}
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
