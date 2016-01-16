package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestContentCodeToInt(t *testing.T) {
	i := contentCodeToInt("mccr")
	if i != 1835230066 {
		t.Errorf("wrong integer value for content code: %v", i)
	}
}

func TestIntToByteArray(t *testing.T) {
	b := intToByteArray(1835230066)
	if !bytes.Equal(b, []byte{109, 99, 99, 114}) {
		t.Errorf("wrong byte array value for int: %v", b)
	}
}

func TestGetServerInfo(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	serverInfoHandler(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	serverHeader := resp.HeaderMap.Get("DAAP-Server")
	if serverHeader == "" {
		t.Error("did not contain DAAP-Server header")
	} else if !strings.Contains(serverHeader, `daap-server`) {
		t.Errorf("DAAP-Server header doesn't match:\n%s", serverHeader)
	}
	if !strings.Contains(string(p), `hello, world`) {
		t.Errorf("response body doen't match:\n%s", p)
	}
}

func TestGetContentCodes(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	contentCodesHandler(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	serverHeader := resp.HeaderMap.Get("DAAP-Server")
	if serverHeader == "" {
		t.Error("did not contain DAAP-Server header")
	} else if !strings.Contains(serverHeader, `daap-server`) {
		t.Errorf("DAAP-Server header doesn't match:\n%s", serverHeader)
	}
	expectedData := []byte{109, 99, 99, 114, 0, 0, 0, 12, 109, 115, 116, 116, 0, 0, 0, 4, 0, 0, 0, 200}
	if !bytes.Equal(p, expectedData) {
		t.Errorf("response body doen't match:\n%v", p)
	}
}
