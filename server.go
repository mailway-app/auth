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

	http.HandleFunc("/", server)
	addr := fmt.Sprintf("127.0.0.1:%d", config.CurrConfig.PortAuth)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func server(w http.ResponseWriter, r *http.Request) {
	method := r.Header.Get("Auth-Method")

	if method == "none" {
		w.Header().Set("Auth-Status", "OK")
		w.Header().Set("Auth-Server", "127.0.0.1")
		w.Header().Set("Auth-Port", fmt.Sprintf("%d", config.CurrConfig.PortForwarding))
		w.WriteHeader(200)
		w.Write([]byte(""))
		return
	}

	if method == "plain" {
		user := r.Header.Get("Auth-User")
		pass := r.Header.Get("Auth-Pass")

		ok, err := login(user, pass)
		if err != nil {
			log.Errorf("failed to login: %s", err)
			w.Header().Set("Auth-Status", "Error during login")
		} else {
			if ok {
				w.Header().Set("Auth-Status", "OK")
				w.Header().Set("Auth-Server", "127.0.0.1")
				w.Header().Set("Auth-Port", fmt.Sprintf("%d", config.CurrConfig.PortResponder))
			} else {
				w.Header().Set("Auth-Status", "Invalid login or password")
			}
		}
		w.WriteHeader(200)
		w.Write([]byte(""))
		return
	}

	w.Header().Set("Auth-Status", "Login method not supported")
	w.WriteHeader(200)
	w.Write([]byte(""))
}
