package mocking

import "net/http"

type Caller interface {
	Call(url string) (*http.Response, error)
}

type caller struct {
	cl *http.Client
}

func NewCaller(cl *http.Client) *caller {
	return &caller{cl: cl}
}

func (c *caller) Call(url string) (*http.Response, error) {
	return c.cl.Get(url)
}
