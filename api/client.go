package api

import (
	"fmt"
	"log"
	"net/http"
)

func Config(storeName string) {
	// todo: load from config
	// host := HOSTNAME
	// port := PORT

	uri := fmt.Sprintf("http://localhost:8080?configure=%s", storeName)
	resp, err := http.Get(uri)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Response status:", resp.Status)
	defer resp.Body.Close()
}
