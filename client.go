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

type SimpleHttpClient struct {
	client      *http.Client
	username    string
	password    string
	initialized bool
}

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

func (s *SimpleHttpClient) SetAuthentication(u, p string) {
	s.username = u
	s.password = p
}
