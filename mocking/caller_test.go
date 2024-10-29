package mocking_test

import (
	"net/http"
	"testing"

	"github.com/aaronschweig/wwi24sea-testing-example/mocking"
)

type testStruct struct{}

func (t *testStruct) Get(url string) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
	}, nil
}

func TestCaller(t *testing.T) {

	caller := mocking.NewCaller(&testStruct{})

	res, err := caller.Call("https://google.com")
	if err != nil {
		t.Fail()
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal("expected status to be ok")
	}
}
