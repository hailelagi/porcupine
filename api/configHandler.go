package api

import (
	"fmt"
	"html/template"
	"net/http"
)

func ConfigHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tmpl, err := template.ParseFiles("./assets/config.html")
		info := HomeInfo{"HashMap", false, []string{}, "8080"}

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err := tmpl.Execute(w, info); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
	case "POST":

		fmt.Fprintf(w, "TODO: %s", r.URL.Path)
	default:
		fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
	}

}
