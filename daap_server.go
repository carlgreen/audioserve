package main

import (
	"fmt"
	"net/http"
)

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

func serverInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)
	fmt.Fprintf(w, "hello, world\n")
}

func contentCodesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)

	data := []byte{}
	data = append(data, "mccr"...)
	data = append(data, intToByteArray(134)...)

	data = append(data, "mstt"...)
	data = append(data, intToData(200)...)

	data = append(data, "mdcl"...)
	data = append(data, intToByteArray(12+31+10)...)
	data = append(data, "mcnm"...)
	data = append(data, stringToData("abal")...)

	data = append(data, "mcna"...)
	data = append(data, stringToData("daap.browsealbumlisting")...)

	data = append(data, "mcty"...)
	data = append(data, shortToData(12)...) // container

	data = append(data, "mdcl"...)
	data = append(data, intToByteArray(12+31+10)...)
	data = append(data, "mcnm"...)
	data = append(data, stringToData("msrv")...)

	data = append(data, "mcna"...)
	data = append(data, stringToData("dmap.serverinforesponse")...)

	data = append(data, "mcty"...)
	data = append(data, shortToData(12)...) // container

	w.Write(data)
}

func main() {
	http.HandleFunc("/server-info", serverInfoHandler)
	http.HandleFunc("/content-codes", contentCodesHandler)
	http.ListenAndServe(":3689", nil)
}
