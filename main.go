package main

import (
	"1_helloapi/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	//routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/hello", handlers.HelloHandler)
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/time", handlers.TimeHandler)

	fmt.Println("Server running on http://localhost:8080")

	err := http.ListenAndServe(":8080", logRequest(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
