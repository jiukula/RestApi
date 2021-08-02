package main

import (
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
