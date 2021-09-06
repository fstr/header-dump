package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	log.Println("Starting header-dump server ...")

	httpListenPort := os.Getenv("PORT")
	if httpListenPort == "" {
		httpListenPort = "8080"
	}

	hostPort := net.JoinHostPort("0.0.0.0", httpListenPort)

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		return
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Save a copy of this request for debugging.
		requestDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintf(w, string(requestDump))
	})

	s := &http.Server{
		Addr:    hostPort,
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
