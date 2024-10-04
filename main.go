package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/niconiahi/one-piece-on-stone/routes/home"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		print("incoming request from GET /\n")
		h := home.Handler{}
		t := template.Must(h.GetTemplate())
		t.Execute(w, h.GetData())
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
