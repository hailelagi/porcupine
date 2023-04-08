package api

import (
	"fmt"
	"net/http"
)

func Serve() {
	// TODO: PROBLEM UNRESOLVED - HOW TO CONNECT THE CLI TO A LONG RUNNING PROCESS IN-MEMORY
	// AND PROVIDE RUNTIME INTROSPECTION via a LOG stream of running operations

	// IDEA1: Expose a global porcupine.Store
	// IDEA2: FIND A WAY TO DO sys IPC CALLS, either by fork() on the current cli process?
	// IDEA3: use a third party tool to persist in-memory state or write to disk...

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
