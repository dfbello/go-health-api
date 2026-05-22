package main

import (
	"net/http"
	"fmt"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8" )
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"status\": \"ok\"}"))
	}

func main() {
	
	sMux := http.NewServeMux()

	server := http.Server{
		Handler: sMux,
		Addr: ":8080",
	}
	
	sMux.HandleFunc("GET /health", healthHandler)

	sMux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8" )
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{\"message\": \"Health API is up and running\"}"))
	})

	fmt.Printf("[INFO] Server listening at %s\n", server.Addr)
	server.ListenAndServe()
}
