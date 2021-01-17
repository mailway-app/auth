package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", AuthServer)
	if err := http.ListenAndServe("127.0.0.1:9000", nil); err != nil {
		panic(err)
	}
}

func AuthServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auth request")
	w.Header().Set("Auth-Status", "OK")
	w.Header().Set("Auth-Server", "127.0.0.1")
	w.Header().Set("Auth-Port", "2500")
	w.WriteHeader(200)
	w.Write([]byte(""))
}
