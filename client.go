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

type simpleHttpClient struct {
	client      *http.Client
	username    string
	password    string
	initialized bool
}

func NewClient() *simpleHttpClient {
	c := &simpleHttpClient{}
	return c
}

func (s *simpleHttpClient) init() {
	tr := &http.Transport{
		//TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	s.client = &http.Client{Transport: tr}
}

func (s *simpleHttpClient) SetAuthentication(u, p string) {
	s.username = u
	s.password = p
}
