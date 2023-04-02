package porcupine

import (
	"fmt"
	"net/http"
)

func Serve() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
