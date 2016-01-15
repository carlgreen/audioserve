package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	w := httptest.NewRecorder()
	handler(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, w.Code)
	}
}
