package api

import (
	"fmt"
	"log"
	"net/http"
)

func Config(storeName string) {
	uri := fmt.Sprintf("http://localhost:8080?configure=%s", storeName)
	resp, err := http.Get(uri)

	if err != nil {
		log.Fatalf("Failed to configure store %s: %v", storeName, err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
