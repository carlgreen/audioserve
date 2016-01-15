package main

import (
	"fmt"
	"net/http"
)

func serverInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(`DAAP-Server`, `daap-server: 1.0`)
	fmt.Fprintf(w, "hello, world\n")
}

func main() {
	http.HandleFunc("/server-info", serverInfoHandler)
	http.ListenAndServe(":3689", nil)
}
