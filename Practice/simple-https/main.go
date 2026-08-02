package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello, Secure World! You are connected via HTTPS.\n"))
	})

	log.Println("Starting HTTPS server on https://localhost:8443")

	err := http.ListenAndServeTLS("localhost:8443", "cert.pem", "key.pem", mux)
	if err != nil {
		log.Fatalf("Server failed to start (do you have the cert.pem and key.pem?): %v", err)
	}
}
