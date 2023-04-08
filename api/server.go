package api

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Serve() {
	// TODO: PROBLEM UNRESOLVED - HOW TO CONNECT THE CLI TO A LONG RUNNING PROCESS IN-MEMORY
	// AND PROVIDE RUNTIME INTROSPECTION via a LOG stream of running operations

	// IDEA1: EXPOSE A main.go SERVER pass around data in a context or expose a porcupine.Store
	// IDEA2: FIND A WAY TO DO sys IPC CALLS, either by fork() on the current cli process?
	// IDEA3: use a third party tool to persist in-memory state or write to disk...

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Listening you are on: %s", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}

// https://go.dev/blog/context
// https://pkg.go.dev/context
func handleSearch(w http.ResponseWriter, req *http.Request) {
	// ctx is the Context for this handler. Calling cancel closes the
	// ctx.Done channel, which is the cancellation signal for requests
	// started by this handler.
	var (
		ctx    context.Context
		cancel context.CancelFunc
	)
	timeout, err := time.ParseDuration(req.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context that is
		// canceled automatically when the timeout expires.
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel() // Cancel ctx as soon as handleSearch returns.
}
