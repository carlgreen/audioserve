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

const DmapShort int16 = 3
const DmapLongLong int16 = 5
const DmapString int16 = 9
const DmapContainer int16 = 12

var contentCodes = []ContentCode{
	{"mstt", "dmap.status", DmapLongLong},
	{"mdcl", "dmap.dictionary", DmapContainer},
	{"msrv", "dmap.serverinforesponse", DmapContainer},
	{"mccr", "dmap.contentcodesresponse", DmapContainer},
	{"mcnm", "dmap.contentcodesnumber", DmapLongLong},
	{"mcna", "dmap.contentcodesname", DmapString},
	{"mcty", "dmap.contentcodestype", DmapShort},
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
	data := [4]byte{
		byte((version.major >> 8) & 0xFF),
		byte(version.major & 0xFF),
		byte(version.minor & 0xFF),
		byte(version.patch & 0xFF),
	}
	return data[:]
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

	data := []byte{}
	data = append(data, "msrv"...)
	data = append(data, intToByteArray(8+8)...)

	data = append(data, "mstt"...)
	data = append(data, intToData(200)...)

	w.Write(data)
}

func contentCodesHandler(contentCodes []ContentCode) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)

		data := []byte{}
		data = append(data, "mccr"...)
		data = append(data, intToByteArray(134)...)

		data = append(data, "mstt"...)
		data = append(data, intToData(200)...)

		for _, contentCode := range contentCodes {
			data = append(data, contentCodeToData(contentCode)...)
		}

		w.Write(data)
	})
}

func main() {
	http.HandleFunc("/server-info", serverInfoHandler)
	http.HandleFunc("/content-codes", contentCodesHandler(contentCodes))
	http.ListenAndServe(":3689", nil)
}
