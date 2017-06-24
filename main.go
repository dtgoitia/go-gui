package main

import (
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/toqueteos/webbrowser"
)

func main() {
	// Create a new rounter
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))

	// Open the app into the default browser
	webbrowser.Open("http://localhost:12345")

	log.Fatal(http.ListenAndServe(":12345", router))
}
