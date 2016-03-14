package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/husobee/vestigo"
)

func headers(inner func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s", r.Method, r.RequestURI)

		w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)

		inner(w, r)
	})
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, r.RequestURI+" not found", http.StatusNotFound)
}

func serverInfoHandler(w http.ResponseWriter, r *http.Request) {
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

func databasesHandler(databases []Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerData := []byte("avdb")

		data := []byte{}

		data = append(data, "mstt"...)
		data = append(data, intToData(200)...)

		data = append(data, "muty"...)
		data = append(data, charToData(0)...)

		data = append(data, "mtco"...)
		data = append(data, intToData(len(databases))...)

		data = append(data, "mrco"...)
		data = append(data, intToData(len(databases))...)

		listing := []byte{}
		for _, database := range databases {
			listing = append(listing, databaseToData(database)...)
		}

		data = append(data, "mlcl"...)
		data = append(data, intToByteArray(len(listing))...)
		data = append(data, listing...)

		headerData = append(headerData, intToByteArray(len(data))...)
		data = append(headerData, data...)

		w.Write(data)
	})
}

func databaseItemsHandler(databases []Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		itemIdParam := vestigo.Param(r, "itemId")
		dbId, err := strconv.Atoi(itemIdParam)
		if err != nil {
			msg := fmt.Sprintf("Cannot convert '%v' to int", itemIdParam)
			log.Print(msg)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

		// TODO do something with this
		// log.Println("meta:", r.Form["meta"])

		// TODO error check this
		database := databases[dbId-1]

		headerData := []byte("adbs")

		data := []byte{}

		data = append(data, "mstt"...)
		data = append(data, intToData(200)...)

		data = append(data, "muty"...)
		data = append(data, charToData(0)...)

		data = append(data, "mtco"...)
		data = append(data, intToData(len(database.songs))...)

		data = append(data, "mrco"...)
		data = append(data, intToData(len(database.songs))...)

		listing := []byte{}
		for _, song := range database.songs {
			listing = append(listing, songToData(song)...)
		}

		data = append(data, "mlcl"...)
		data = append(data, intToByteArray(len(listing))...)
		data = append(data, listing...)

		headerData = append(headerData, intToByteArray(len(data))...)
		data = append(headerData, data...)

		w.Write(data)
	})
}

func databaseContainersHandler(databases []Database) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerData := []byte("aply")

		data := []byte{}

		data = append(data, "mstt"...)
		data = append(data, intToData(200)...)

		data = append(data, "muty"...)
		data = append(data, charToData(0)...)

		data = append(data, "mtco"...)
		data = append(data, intToData(0)...)

		data = append(data, "mrco"...)
		data = append(data, intToData(0)...)

		headerData = append(headerData, intToByteArray(len(data))...)
		data = append(headerData, data...)

		w.Write(data)
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	headerData := []byte("mlog")

	data := []byte{}

	data = append(data, "mstt"...)
	data = append(data, intToData(200)...)

	data = append(data, "mlid"...)
	// TODO generate a real session ID
	data = append(data, intToData(113)...)

	headerData = append(headerData, intToByteArray(len(data))...)
	data = append(headerData, data...)

	w.Write(data)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// TODO end session r.URL.Query().Get("session-id")
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	revNumParam := r.Form.Get("revision-number")
	revNum, err := strconv.Atoi(revNumParam)
	if err != nil {
		msg := fmt.Sprintf("Cannot convert '%v' to int", revNumParam)
		log.Print(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	headerData := []byte("mupd")

	data := []byte{}

	data = append(data, "musr"...)
	data = append(data, intToData(revNum)...)

	data = append(data, "mstt"...)
	data = append(data, intToData(200)...)

	headerData = append(headerData, intToByteArray(len(data))...)
	data = append(headerData, data...)

	w.Write(data)
}
