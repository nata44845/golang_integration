package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	exp := "message from handler"
	srv := httptest.NewServer(http.HandlerFunc(handler))
	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	textBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer resp.Body.Close()

	text := string(textBytes)
	if text != exp {
		t.Logf("%v не равно %v", text, exp)
		t.Fail()
	}
}
