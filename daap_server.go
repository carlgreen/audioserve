package main

import (
	"fmt"
	"net/http"
)

func serverInfoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello, world\n")
}

func main() {
	http.HandleFunc("/server-info", serverInfoHandler)
	http.ListenAndServe(":3689", nil)
}
