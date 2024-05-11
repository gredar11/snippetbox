package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "IP address")

	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet", showSnippet)

	mux.HandleFunc("/snippet/create", createSnippet)

	log.Printf("Запус веб-сервера на %s", *addr)

	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
