package main

import (
	"fmt"
	"net/http"

	"github.com/mailway-app/config"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := config.Init(); err != nil {
		log.Fatalf("failed to init config: %s", err)
	}

	http.HandleFunc("/", AuthServer)
	addr := fmt.Sprintf("127.0.0.1:%d", config.CurrConfig.PortAuth)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func AuthServer(w http.ResponseWriter, r *http.Request) {
	log.Debugf("auth request")
	w.Header().Set("Auth-Status", "OK")
	w.Header().Set("Auth-Server", "127.0.0.1")
	w.Header().Set("Auth-Port", fmt.Sprintf("%d", config.CurrConfig.PortForwarding))
	w.WriteHeader(200)
	w.Write([]byte(""))
}
