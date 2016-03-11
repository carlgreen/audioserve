package main

import (
	"log"
	"net/http"

	"github.com/husobee/vestigo"
)

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
	{"mupd", "dmap.updateresponse", DmapContainer},
	{"musr", "dmap.serverrevision", DmapLong},
	{"muty", "dmap.updatetype", DmapChar},
	{"mccr", "dmap.contentcodesresponse", DmapContainer},
	{"mcnm", "dmap.contentcodesnumber", DmapLong},
	{"mcna", "dmap.contentcodesname", DmapString},
	{"mcty", "dmap.contentcodestype", DmapShort},
	{"apro", "daap.protocolversion", DmapVersion},
	{"avdb", "daap.serverdatabases", DmapContainer},
	{"adbs", "daap.databasesongs", DmapContainer},
}

var databases = []Database{
	{
		ListingItem{1, 1, "testdb", 0, 0},
		[]ListingItem{
			{1, 1, "some item", 0, 0},
		},
		nil,
	},
}

func routes(contentCodes []ContentCode, databases []Database) http.Handler {
	router := vestigo.NewRouter()
	router.Get("/server-info", headers(serverInfoHandler))
	router.Get("/content-codes", headers(contentCodesHandler(contentCodes)))
	router.Get("/databases", headers(databasesHandler(databases)))
	router.Get("/databases/:itemId/items", headers(databaseItemsHandler(databases)))
	router.Get("/login", headers(loginHandler))
	router.Get("/logout", headers(logoutHandler))
	router.Get("/update", headers(updateHandler))
	vestigo.CustomNotFoundHandlerFunc(headers(defaultHandler))
	return router
}

func main() {
	router := routes(contentCodes, databases)
	log.Fatal(http.ListenAndServe(":3689", router))
}
