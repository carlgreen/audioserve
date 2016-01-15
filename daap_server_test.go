package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	serverInfoHandler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, w.Code)
	}
	p, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(p), `hello, world`) {
		t.Errorf("response body doen't match:\n%s", p)
	}
}
