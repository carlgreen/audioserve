package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestIntToByteArray(t *testing.T) {
	b := intToByteArray(1835230066)
	if !bytes.Equal(b, []byte{109, 99, 99, 114}) {
		t.Errorf("wrong byte array value for int: %v", b)
	}
}

func TestShortToByteArray(t *testing.T) {
	b := shortToByteArray(12)
	if !bytes.Equal(b, []byte{0, 12}) {
		t.Errorf("wrong byte array value for short: %v", b)
	}
}

func TestStringToData(t *testing.T) {
	data := stringToData("abcde")
	if !bytes.Equal(data, []byte{0, 0, 0, 5, 97, 98, 99, 100, 101}) {
		t.Errorf("wrong byte array value for string: %v", data)
	}
}

func TestIntToData(t *testing.T) {
	data := intToData(200)
	if !bytes.Equal(data, []byte{0, 0, 0, 4, 0, 0, 0, 200}) {
		t.Errorf("wrong byte array value for int: %v", data)
	}
}

func TestShortToData(t *testing.T) {
	data := shortToData(12)
	if !bytes.Equal(data, []byte{0, 0, 0, 2, 0, 12}) {
		t.Errorf("wrong byte array value for short: %v", data)
	}
}

func TestContentCodeToData(t *testing.T) {
	data := contentCodeToData(ContentCode{"abal", "daap.browsealbumlisting", 12})
	expectedData := []byte{
		109, 100, 99, 108, 0, 0, 0, 53, // mdcl
		109, 99, 110, 109, 0, 0, 0, 4, 97, 98, 97, 108, // mcnm (abal)
		109, 99, 110, 97, 0, 0, 0, 23, 100, 97, 97, 112, 46, 98, 114, 111, 119, 115, 101, 97, 108, 98, 117, 109, 108, 105, 115, 116, 105, 110, 103, // mcna
		109, 99, 116, 121, 0, 0, 0, 2, 0, 12, // mcty
	}
	if !bytes.Equal(data, expectedData) {
		t.Errorf("wrong byte array value for content code structure: %v", data)
	}
}

func TestVersionToData(t *testing.T) {
	data := versionToData(Version{1, 0, 0})
	if !bytes.Equal(data, []byte{0, 0, 0, 4, 0, 1, 0, 0}) {
		t.Errorf("wrong byte array value for version: %v", data)
	}
}

func TestCharToData(t *testing.T) {
	data := charToData(1)
	if !bytes.Equal(data, []byte{0, 0, 0, 1, 1}) {
		t.Errorf("wrong byte array value for char: %v", data)
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

	expectedData := []byte{
		109, 115, 114, 118, 0, 0, 0, 160, // msrv
		109, 115, 116, 116, 0, 0, 0, 4, 0, 0, 0, 200, // mstt
		109, 112, 114, 111, 0, 0, 0, 4, 0, 1, 0, 0, // mpro
		97, 112, 114, 111, 0, 0, 0, 4, 0, 1, 0, 0, // apro
		109, 105, 110, 109, 0, 0, 0, 11, 100, 97, 97, 112, 45, 115, 101, 114, 118, 101, 114, // minm
		109, 115, 108, 114, 0, 0, 0, 1, 1, // mslr
		109, 115, 116, 109, 0, 0, 0, 4, 0, 0, 7, 8, // mstm
		109, 115, 97, 108, 0, 0, 0, 1, 1, // msal
		109, 115, 117, 112, 0, 0, 0, 1, 1, // msup
		109, 115, 112, 105, 0, 0, 0, 1, 1, // mspi
		109, 115, 101, 120, 0, 0, 0, 1, 1, // msex
		109, 115, 98, 114, 0, 0, 0, 1, 1, // msbr
		109, 115, 113, 121, 0, 0, 0, 1, 1, // msqy
		109, 115, 105, 120, 0, 0, 0, 1, 1, // msix
		109, 115, 114, 115, 0, 0, 0, 1, 1, // msrs
		109, 115, 100, 99, 0, 0, 0, 4, 0, 0, 0, 1, // msdc
	}
	if !bytes.Equal(p, expectedData) {
		t.Errorf("response body doen't match:\n%v", p)
	}
}

func TestGetContentCodes(t *testing.T) {
	var contentCodes = []ContentCode{
		{"abal", "daap.browsealbumlisting", DmapContainer},
		{"msrv", "dmap.serverinforesponse", DmapContainer},
	}
	contentCodesHandle := contentCodesHandler(contentCodes)
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	contentCodesHandle(resp, req)
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

	expectedData := []byte{
		109, 99, 99, 114, 0, 0, 0, 134, // mccr
		109, 115, 116, 116, 0, 0, 0, 4, 0, 0, 0, 200, // mstt
		109, 100, 99, 108, 0, 0, 0, 53, // mdcl
		109, 99, 110, 109, 0, 0, 0, 4, 97, 98, 97, 108, // mcnm (abal)
		109, 99, 110, 97, 0, 0, 0, 23, 100, 97, 97, 112, 46, 98, 114, 111, 119, 115, 101, 97, 108, 98, 117, 109, 108, 105, 115, 116, 105, 110, 103, // mcna
		109, 99, 116, 121, 0, 0, 0, 2, 0, 12, // mcty
		109, 100, 99, 108, 0, 0, 0, 53, // mdcl
		109, 99, 110, 109, 0, 0, 0, 4, 109, 115, 114, 118, // mcnm (msrv)
		109, 99, 110, 97, 0, 0, 0, 23, 100, 109, 97, 112, 46, 115, 101, 114, 118, 101, 114, 105, 110, 102, 111, 114, 101, 115, 112, 111, 110, 115, 101, // mcna
		109, 99, 116, 121, 0, 0, 0, 2, 0, 12, // mcty
	}
	if !bytes.Equal(p, expectedData) {
		t.Errorf("response body doen't match:\n%v", p)
	}
}

func TestGetLogin(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	loginHandler(resp, req)
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

	expectedData := []byte{
		109, 108, 111, 103, 0, 0, 0, 24, // mlog
		109, 115, 116, 116, 0, 0, 0, 4, 0, 0, 0, 200, // mstt
		109, 108, 105, 100, 0, 0, 0, 4, 0, 0, 0, 113, // mlid
	}
	if !bytes.Equal(p, expectedData) {
		t.Errorf("response body doen't match:\n%v", p)
	}
}

func TestGetLogout(t *testing.T) {
	req, err := http.NewRequest("GET", "?session-id=113", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	logoutHandler(resp, req)
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

	if !bytes.Equal(p, []byte{}) {
		t.Errorf("response body doen't match:\n%v", p)
	}
}
