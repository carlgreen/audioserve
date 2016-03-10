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
	items          []ListingItem
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
	{"adbs", "daap.databasesongs", DmapContainer},
}

var databases = []ListingItem{
	{1, 1, "testdb", 0, 0, []ListingItem{
		{1, 1, "some item", 0, 0, nil},
	}},
}

func routes(contentCodes []ContentCode, databases []ListingItem) http.Handler {
	router := vestigo.NewRouter()
	router.Get("/server-info", headers(serverInfoHandler))
	router.Get("/content-codes", headers(contentCodesHandler(contentCodes)))
	router.Get("/databases", headers(databasesHandler(databases)))
	router.Get("/databases/:itemId/items", headers(databaseItemsHandler(databases)))
	router.Get("/login", headers(loginHandler))
	router.Get("/logout", headers(logoutHandler))
	return router
}

func main() {
	router := routes(contentCodes, databases)
	log.Fatal(http.ListenAndServe(":3689", router))
}
