package main

import (
	"net/http"
)

type ContentCode struct {
	number   string
	name     string
	dmapType int16
}

type Version struct {
	major uint16
	minor uint8
	patch uint8
}

const DmapChar int16 = 1
const DmapShort int16 = 3
const DmapLong int16 = 5
const DmapString int16 = 9
const DmapVersion int16 = 11
const DmapContainer int16 = 12

var contentCodes = []ContentCode{
	{"minm", "dmap.itemname", DmapString},
	{"mstt", "dmap.status", DmapLong},
	{"mdcl", "dmap.dictionary", DmapContainer},
	{"msrv", "dmap.serverinforesponse", DmapContainer},
	{"mslr", "dmap.loginrequired", DmapChar},
	{"mpro", "dmap.protocolversion", DmapVersion},
	{"msal", "dmap.supportsautologout", DmapChar},
	{"msup", "dmap.supportsupdate", DmapChar},
	{"mspi", "dmap.supportspersistentids", DmapChar},
	{"msex", "dmap.supportsextensions", DmapChar},
	{"msbr", "dmap.supportsbrowse", DmapChar},
	{"msqy", "dmap.supportsquery", DmapChar},
	{"msix", "dmap.supportsindex", DmapChar},
	{"msrs", "dmap.supportsresolve", DmapChar},
	{"mstm", "dmap.timeoutinterval", DmapLong},
	{"msdc", "dmap.databasescount", DmapLong},
	{"mccr", "dmap.contentcodesresponse", DmapContainer},
	{"mcnm", "dmap.contentcodesnumber", DmapLong},
	{"mcna", "dmap.contentcodesname", DmapString},
	{"mcty", "dmap.contentcodestype", DmapShort},
	{"apro", "daap.protocolversion", DmapVersion},
}

func intToByteArray(i int) []byte {
	data := [4]byte{
		byte((i >> 24) & 0xFF),
		byte((i >> 16) & 0xFF),
		byte((i >> 8) & 0xFF),
		byte(i & 0xFF),
	}
	return data[:]
}

func shortToByteArray(i int16) []byte {
	data := [2]byte{
		byte((i >> 8) & 0xFF),
		byte(i & 0xFF),
	}
	return data[:]
}

func stringToData(s string) []byte {
	data := intToByteArray(len(s))
	data = append(data, s...)
	return data
}

func intToData(i int) []byte {
	data := intToByteArray(4)
	data = append(data, intToByteArray(i)...)
	return data
}

func shortToData(i int16) []byte {
	data := intToByteArray(2)
	data = append(data, shortToByteArray(i)...)
	return data
}

func versionToData(version Version) []byte {
	data := intToByteArray(4)
	versionData := [4]byte{
		byte((version.major >> 8) & 0xFF),
		byte(version.major & 0xFF),
		byte(version.minor & 0xFF),
		byte(version.patch & 0xFF),
	}
	data = append(data, versionData[:]...)
	return data
}

func charToData(b byte) []byte {
	data := intToByteArray(1)
	data = append(data, b)
	return data
}

func contentCodeToData(contentCode ContentCode) []byte {
	data := []byte{}

	data = append(data, "mdcl"...)
	data = append(data, intToByteArray(12+8+len(contentCode.name)+10)...)

	data = append(data, "mcnm"...)
	data = append(data, stringToData(contentCode.number)...)

	data = append(data, "mcna"...)
	data = append(data, stringToData(contentCode.name)...)

	data = append(data, "mcty"...)
	data = append(data, shortToData(contentCode.dmapType)...)

	return data
}

func serverInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)

	headerData := []byte("msrv")

	data := []byte{}

	data = append(data, "mstt"...)
	data = append(data, intToData(200)...)

	data = append(data, "mpro"...)
	data = append(data, versionToData(Version{1, 0, 0})...)

	data = append(data, "apro"...)
	data = append(data, versionToData(Version{1, 0, 0})...)

	data = append(data, "minm"...)
	data = append(data, stringToData("daap-server")...)

	data = append(data, "mslr"...)
	data = append(data, charToData(1)...)

	data = append(data, "mstm"...)
	data = append(data, intToData(1800)...)

	data = append(data, "msal"...)
	data = append(data, charToData(1)...)

	data = append(data, "msup"...)
	data = append(data, charToData(1)...)

	data = append(data, "mspi"...)
	data = append(data, charToData(1)...)

	data = append(data, "msex"...)
	data = append(data, charToData(1)...)

	data = append(data, "msbr"...)
	data = append(data, charToData(1)...)

	data = append(data, "msqy"...)
	data = append(data, charToData(1)...)

	data = append(data, "msix"...)
	data = append(data, charToData(1)...)

	data = append(data, "msrs"...)
	data = append(data, charToData(1)...)

	data = append(data, "msdc"...)
	data = append(data, intToData(1)...)

	headerData = append(headerData, intToByteArray(len(data))...)
	data = append(headerData, data...)

	w.Write(data)
}

func contentCodesHandler(contentCodes []ContentCode) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)

		headerData := []byte("mccr")

		data := []byte{}

		data = append(data, "mstt"...)
		data = append(data, intToData(200)...)

		for _, contentCode := range contentCodes {
			data = append(data, contentCodeToData(contentCode)...)
		}

		headerData = append(headerData, intToByteArray(len(data))...)
		data = append(headerData, data...)

		w.Write(data)
	})
}

func main() {
	http.HandleFunc("/server-info", serverInfoHandler)
	http.HandleFunc("/content-codes", contentCodesHandler(contentCodes))
	http.ListenAndServe(":3689", nil)
}
