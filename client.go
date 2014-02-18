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
}

// NewClient instantiates a pointer to a SimpleHttpClient object, which is
// the base of all client operations.
func NewClient() *SimpleHttpClient {
	c := &SimpleHttpClient{}
	return c
}

func (s *SimpleHttpClient) init() {
	tr := &http.Transport{
		//TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	s.client = &http.Client{Transport: tr}
}

// SetAuthentication sets BASIC authentication username and password for
// all further client requests.
func (s *SimpleHttpClient) SetAuthentication(u, p string) {
	s.username = u
	s.password = p
}
