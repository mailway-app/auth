package main

import (
	"fmt"
	"net/http"

	mailwayConfig "github.com/mailway-app/config"
)

var config *mailwayConfig.Config

func main() {
	var err error
	config, err = mailwayConfig.Read()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", AuthServer)
	addr := fmt.Sprintf("127.0.0.1:%d", config.PortAuth)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func AuthServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("auth request")
	w.Header().Set("Auth-Status", "OK")
	w.Header().Set("Auth-Server", "127.0.0.1")
	w.Header().Set("Auth-Port", fmt.Sprintf("%d", config.PortForwarding))
	w.WriteHeader(200)
	w.Write([]byte(""))
}
