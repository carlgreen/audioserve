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
	if !bytes.Equal(data, []byte{0, 1, 0, 0}) {
		t.Errorf("wrong byte array value for version: %v", data)
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
