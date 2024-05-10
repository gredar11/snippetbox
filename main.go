package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Форма показа заметки"))
	})

	mux.HandleFunc("/snippet/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Форма создания заметки"))
	})

	log.Println("Запус веб-сервера на http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Привет из SnippetBox"))
}
