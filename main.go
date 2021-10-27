package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/mario", marioHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hai Bang, Belajar bikin Web pake golang"))
}

func marioHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Route Mario Page"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Masuk ke Home"))
}
