package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

var listenPort int

func init() {
	flag.IntVar(&listenPort, "p", 1234, "Port to listen on")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			log.Println(err)
		}

		w.WriteHeader(http.StatusOK)

		fmt.Printf(">>> %s\n", ip)
		fmt.Println(string(dump))
	})

	log.Printf("--- Listening on port :%d\n", listenPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil); err != nil {
		log.Fatalln(err)
	}
}
