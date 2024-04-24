package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type HomeInfo struct {
	Porcupine   string
	Clustering  bool
	Nodes       []string
	DefaultPort string
}

func Serve() {
	http.HandleFunc("/configure", ConfigHandler)
	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/ls", lsHandler)
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/rm", rmHandler)
	http.HandleFunc("/line", plotHandler)
	http.HandleFunc("/", homeHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
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

func plotHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./assets/line.html")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, nil)
	fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./assets/index.html")
	//info := HomeInfo{"HashMap", true, []string{"8080", "8081"}}
	info := HomeInfo{"HashMap", false, []string{}, "8080"}

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	tmpl.Execute(w, info)
}
