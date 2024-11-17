package mocking

import "net/http"

type Caller interface {
	Call(url string) (*http.Response, error)
}

type Getter interface {
	Get(url string) (*http.Response, error)
}

type caller struct {
	cl Getter
}

func NewCaller(cl Getter) *caller {
	return &caller{cl: cl}
}

func (c *caller) Call(url string) (*http.Response, error) {
	return c.cl.Get(url)
}
