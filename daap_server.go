package main

import (
	"log"
	"net/http"

	"github.com/husobee/vestigo"
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

type ListingItem struct {
	itemId         int   // uint32
	persistentId   int64 // uint64
	itemName       string
	itemCount      int // uint32
	containerCount int // uint32
}

const DmapChar int16 = 1
const DmapShort int16 = 3
const DmapLong int16 = 5
const DmapLongLong int16 = 7
const DmapString int16 = 9
const DmapVersion int16 = 11
const DmapContainer int16 = 12

var contentCodes = []ContentCode{
	{"miid", "dmap.itemid", DmapLong},
	{"minm", "dmap.itemname", DmapString},
	{"mper", "dmap.persistentid", DmapLongLong},
	{"mstt", "dmap.status", DmapLong},
	{"mimc", "dmap.itemcount", DmapLong},
	{"mctc", "dmap.containercount", DmapLong},
	{"mrco", "dmap.returnedcount", DmapLong},
	{"mtco", "dmap.specifiedtotalcount", DmapLong},
	{"mlcl", "dmap.listing", DmapContainer},
	{"mlit", "dmap.listingitem", DmapContainer},
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
	{"mlog", "dmap.loginresponse", DmapContainer},
	{"mlid", "dmap.sessionid", DmapLong},
	{"muty", "dmap.updatetype", DmapChar},
	{"mccr", "dmap.contentcodesresponse", DmapContainer},
	{"mcnm", "dmap.contentcodesnumber", DmapLong},
	{"mcna", "dmap.contentcodesname", DmapString},
	{"mcty", "dmap.contentcodestype", DmapShort},
	{"apro", "daap.protocolversion", DmapVersion},
	{"avdb", "daap.serverdatabases", DmapContainer},
}

var databases = []ListingItem{
	{1, 1, "testdb", 0, 0},
}

func longToByteArray(l int64) []byte {
	data := [8]byte{
		byte((l >> 56) & 0xFF),
		byte((l >> 48) & 0xFF),
		byte((l >> 40) & 0xFF),
		byte((l >> 32) & 0xFF),
		byte((l >> 24) & 0xFF),
		byte((l >> 16) & 0xFF),
		byte((l >> 8) & 0xFF),
		byte(l & 0xFF),
	}
	return data[:]
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

func longToData(i int64) []byte {
	data := intToByteArray(8)
	data = append(data, longToByteArray(i)...)
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

func listingItemToData(listingItem ListingItem) []byte {
	data := []byte{}

	data = append(data, "mlit"...)
	data = append(data, intToByteArray(12+16+8+len(listingItem.itemName)+12+12)...)

	data = append(data, "miid"...)
	data = append(data, intToData(listingItem.itemId)...)

	data = append(data, "mper"...)
	data = append(data, longToData(listingItem.persistentId)...)

	data = append(data, "minm"...)
	data = append(data, stringToData(listingItem.itemName)...)

	data = append(data, "mimc"...)
	data = append(data, intToData(listingItem.itemCount)...)

	data = append(data, "mctc"...)
	data = append(data, intToData(listingItem.containerCount)...)

	return data
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

func databasesHandler(databases []ListingItem) http.HandlerFunc {
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
			listing = append(listing, listingItemToData(database)...)
		}

		data = append(data, "mlcl"...)
		data = append(data, intToByteArray(len(listing))...)
		data = append(data, listing...)

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

func headers(inner func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s\t%s", r.Method, r.RequestURI)

		w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)

		inner(w, r)
	})
}

func main() {
	router := vestigo.NewRouter()
	router.Get("/server-info", headers(serverInfoHandler))
	router.Get("/content-codes", headers(contentCodesHandler(contentCodes)))
	router.Get("/databases", headers(databasesHandler(databases)))
	router.Get("/login", headers(loginHandler))
	router.Get("/logout", headers(logoutHandler))
	log.Fatal(http.ListenAndServe(":3689", router))
}
