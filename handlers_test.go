package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetServerInfo(t *testing.T) {
	router := routes(nil, nil)
	req, err := http.NewRequest("GET", "/server-info", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
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
	router := routes(contentCodes, nil)
	req, err := http.NewRequest("GET", "/content-codes", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
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
	router := routes(nil, nil)
	req, err := http.NewRequest("GET", "/login", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
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
	router := routes(nil, nil)
	req, err := http.NewRequest("GET", "/logout?session-id=113", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(p, []byte{}) {
		t.Errorf("response body doen't match:\n%v", p)
	}
}

func TestGetDatabases(t *testing.T) {
	var databases = []Database{
		{ListingItem{1, 1, "testdb", 1, 0}, nil, nil},
	}
	router := routes(nil, databases)

	req, err := http.NewRequest("GET", "/databases", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectedData := []byte{
		97, 118, 100, 98, 0, 0, 0, 127, // avdb
		109, 115, 116, 116, 0, 0, 0, 4, 0, 0, 0, 200, // mstt
		109, 117, 116, 121, 0, 0, 0, 1, 0, // muty
		109, 116, 99, 111, 0, 0, 0, 4, 0, 0, 0, 1, // mtco
		109, 114, 99, 111, 0, 0, 0, 4, 0, 0, 0, 1, // mrco
		109, 108, 99, 108, 0, 0, 0, 74, // mlcl
		109, 108, 105, 116, 0, 0, 0, 66, // mlit
		109, 105, 105, 100, 0, 0, 0, 4, 0, 0, 0, 1, // miid
		109, 112, 101, 114, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 1, // mper
		109, 105, 110, 109, 0, 0, 0, 6, 116, 101, 115, 116, 100, 98, // minm
		109, 105, 109, 99, 0, 0, 0, 4, 0, 0, 0, 1, // mimc
		109, 99, 116, 99, 0, 0, 0, 4, 0, 0, 0, 0, // mctc
	}
	if !bytes.Equal(p, expectedData) {
		t.Errorf("response body doen't match:\n%v", p)
	}
}

func TestGetDatabaseItems(t *testing.T) {
	var databases = []Database{
		{
			ListingItem{1, 1, "testdb", 1, 0},
			[]ListingItem{
				{2, 2, "aname", 0, 0},
			},
			nil,
		},
	}

	router := routes(nil, databases)
	req, err := http.NewRequest("GET", "/databases/1/items?session-id=113&meta=dmap.itemid,dmap.itemname,daap.songalbum,daap.songartist,daap.songformat,daap.songtime,daap.songsize,daap.songgenre,daap.songyear,daap.songtracknumber", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectedData := []byte{
		97, 100, 98, 115, 0, 0, 0, 21 + 24 + 8 + 73, // adbs
		109, 115, 116, 116, 0, 0, 0, 4, 0, 0, 0, 200, // mstt
		109, 117, 116, 121, 0, 0, 0, 1, 0, // muty
		109, 116, 99, 111, 0, 0, 0, 4, 0, 0, 0, 1, // mtco
		109, 114, 99, 111, 0, 0, 0, 4, 0, 0, 0, 1, // mrco
		109, 108, 99, 108, 0, 0, 0, 73, // mlcl
		109, 108, 105, 116, 0, 0, 0, 65, // mlit
		109, 105, 105, 100, 0, 0, 0, 4, 0, 0, 0, 2, // miid
		109, 112, 101, 114, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 2, // mper
		109, 105, 110, 109, 0, 0, 0, 5, 97, 110, 97, 109, 101, // minm
		109, 105, 109, 99, 0, 0, 0, 4, 0, 0, 0, 0, // mimc
		109, 99, 116, 99, 0, 0, 0, 4, 0, 0, 0, 0, // mctc
	}
	if !bytes.Equal(p, expectedData) {
		t.Errorf("response body doen't match:\n%v", p)
	}
}

func TestGetUpdate(t *testing.T) {
	router := routes(nil, nil)
	req, err := http.NewRequest("GET", "/update?session-id=113&revision-number=1", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	p, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	expectedData := []byte{
		109, 117, 112, 100, 0, 0, 0, 24, // mupd
		109, 117, 115, 114, 0, 0, 0, 4, 0, 0, 0, 1, // musr
		109, 115, 116, 116, 0, 0, 0, 4, 0, 0, 0, 200, // mstt
	}
	if !bytes.Equal(p, expectedData) {
		t.Errorf("response body doen't match:\n%v\n%v", p, expectedData)
	}
}

func TestHeaders(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp := httptest.NewRecorder()
	dummyHandlerCalled := false
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		dummyHandlerCalled = true
	}
	headers(dummyHandler)(resp, req)
	if resp.Code != http.StatusOK {
		t.Errorf("wrong http status, want %v, got %v", http.StatusOK, resp.Code)
	}
	serverHeader := resp.HeaderMap.Get("DAAP-Server")
	if serverHeader == "" {
		t.Error("did not contain DAAP-Server header")
	} else if !strings.Contains(serverHeader, `daap-server`) {
		t.Errorf("DAAP-Server header doesn't match:\n%s", serverHeader)
	}
	if !dummyHandlerCalled {
		t.Error("did not call inner handler")
	}
}
