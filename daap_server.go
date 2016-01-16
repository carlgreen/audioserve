package main

import (
	"fmt"
	"net/http"
)

func contentCodeToInt(contentCode string) int {
	return (int(contentCode[0]&0xFF) << 24) |
		(int(contentCode[1]&0xFF) << 16) |
		(int(contentCode[2]&0xFF) << 8) |
		int(contentCode[3]&0xFF)
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

func serverInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)
	fmt.Fprintf(w, "hello, world\n")
}

func contentCodesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)

	data := []byte{}
	data = append(data, intToByteArray(contentCodeToInt("mccr"))...)
	data = append(data, intToByteArray(134)...)

	data = append(data, intToByteArray(contentCodeToInt("mstt"))...)
	data = append(data, intToByteArray(4)...)
	data = append(data, intToByteArray(200)...)

	data = append(data, intToByteArray(contentCodeToInt("mdcl"))...)
	data = append(data, intToByteArray(12+31+10)...)
	data = append(data, intToByteArray(contentCodeToInt("mcnm"))...)
	data = append(data, intToByteArray(4)...)
	data = append(data, intToByteArray(contentCodeToInt("abal"))...)

	data = append(data, intToByteArray(contentCodeToInt("mcna"))...)
	data = append(data, intToByteArray(23)...)
	data = append(data, "daap.browsealbumlisting"...)

	data = append(data, intToByteArray(contentCodeToInt("mcty"))...)
	data = append(data, intToByteArray(2)...)
	data = append(data, shortToByteArray(12)...) // container

	data = append(data, intToByteArray(contentCodeToInt("mdcl"))...)
	data = append(data, intToByteArray(12+31+10)...)
	data = append(data, intToByteArray(contentCodeToInt("mcnm"))...)
	data = append(data, intToByteArray(4)...)
	data = append(data, intToByteArray(contentCodeToInt("msrv"))...)

	data = append(data, intToByteArray(contentCodeToInt("mcna"))...)
	data = append(data, intToByteArray(23)...)
	data = append(data, "dmap.serverinforesponse"...)

	data = append(data, intToByteArray(contentCodeToInt("mcty"))...)
	data = append(data, intToByteArray(2)...)
	data = append(data, shortToByteArray(12)...) // container

	w.Write(data)
}

func main() {
	http.HandleFunc("/server-info", serverInfoHandler)
	http.HandleFunc("/content-codes", contentCodesHandler)
	http.ListenAndServe(":3689", nil)
}
