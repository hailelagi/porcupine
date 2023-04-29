package api

import (
	"fmt"
	"net/http"
)

func Serve() {
	// SPAWN a server process seperate from the CLI process.
	// Perform IPC calls
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
