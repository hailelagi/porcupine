package api

import (
	"fmt"
	"log"
	"net/http"
)

func Serve() {
	http.HandleFunc("/configure", configHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/ls", lsHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/rm", rmHandler)
	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
}

func lsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
}

func rmHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[waiting to dispatch]")
	// todo: serve api docs
	fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
}
