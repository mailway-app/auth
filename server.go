package main

import (
	"fmt"
	"net/http"

	mailwayConfig "github.com/mailway-app/config"

	log "github.com/sirupsen/logrus"
)

var config *mailwayConfig.Config

func main() {
	var err error
	config, err = mailwayConfig.Read()
	if err != nil {
		log.Fatal(err)
	}
	log.SetLevel(config.GetLogLevel())
	log.SetFormatter(config.GetLogFormat())

	http.HandleFunc("/", AuthServer)
	addr := fmt.Sprintf("127.0.0.1:%d", config.PortAuth)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
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
