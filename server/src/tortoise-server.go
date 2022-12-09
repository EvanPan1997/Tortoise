package main

import (
	"log"
	"net/http"
)

func main() {
	// following the path, choose a handler
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":5678", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	// code|message|payload(need to encode)
	responseBody := "{\"code\":\"200\",\"message\":\"ok\"}"
	_, err := w.Write([]byte(responseBody))
	if err != nil {
		return
	}
}
