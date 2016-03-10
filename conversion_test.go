package main

import (
	"bytes"
	"testing"
)

func TestShortToByteArray(t *testing.T) {
	b := shortToByteArray(12)
	if !bytes.Equal(b, []byte{0, 12}) {
		t.Errorf("wrong byte array value for short: %v", b)
	}
}

func TestIntToByteArray(t *testing.T) {
	b := intToByteArray(1835230066)
	if !bytes.Equal(b, []byte{109, 99, 99, 114}) {
		t.Errorf("wrong byte array value for int: %v", b)
	}
}

func TestStringToData(t *testing.T) {
	data := stringToData("abcde")
	if !bytes.Equal(data, []byte{0, 0, 0, 5, 97, 98, 99, 100, 101}) {
		t.Errorf("wrong byte array value for string: %v", data)
	}
}

func TestCharToData(t *testing.T) {
	data := charToData(1)
	if !bytes.Equal(data, []byte{0, 0, 0, 1, 1}) {
		t.Errorf("wrong byte array value for char: %v", data)
	}
}

func TestShortToData(t *testing.T) {
	data := shortToData(12)
	if !bytes.Equal(data, []byte{0, 0, 0, 2, 0, 12}) {
		t.Errorf("wrong byte array value for short: %v", data)
	}
}

func TestIntToData(t *testing.T) {
	data := intToData(200)
	if !bytes.Equal(data, []byte{0, 0, 0, 4, 0, 0, 0, 200}) {
		t.Errorf("wrong byte array value for int: %v", data)
	}
}

func TestLongToData(t *testing.T) {
	data := longToData(200)
	if !bytes.Equal(data, []byte{0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 200}) {
		t.Errorf("wrong byte array value for long: %v", data)
	}
}

func TestVersionToData(t *testing.T) {
	data := versionToData(Version{1, 0, 0})
	if !bytes.Equal(data, []byte{0, 0, 0, 4, 0, 1, 0, 0}) {
		t.Errorf("wrong byte array value for version: %v", data)
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

func TestListingItemToData(t *testing.T) {
	data := listingItemToData(ListingItem{1, 1, "testdb", 1, 0, nil})
	expectedData := []byte{
		109, 108, 105, 116, 0, 0, 0, 66, // mlit
		109, 105, 105, 100, 0, 0, 0, 4, 0, 0, 0, 1, // miid
		109, 112, 101, 114, 0, 0, 0, 8, 0, 0, 0, 0, 0, 0, 0, 1, // mper
		109, 105, 110, 109, 0, 0, 0, 6, 116, 101, 115, 116, 100, 98, // minm
		109, 105, 109, 99, 0, 0, 0, 4, 0, 0, 0, 1, // mimc
		109, 99, 116, 99, 0, 0, 0, 4, 0, 0, 0, 0, // mctc
	}
	if !bytes.Equal(data, expectedData) {
		t.Errorf("wrong byte array value for listing item structure: %v", data)
	}
}
