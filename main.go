package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ondc/on_search", handleRequest("on_search"))
	http.HandleFunc("/ondc/on_select", handleRequest("on_select"))
	http.HandleFunc("/ondc/on_init", handleRequest("on_init"))
	http.HandleFunc("/ondc/on_confirm", handleRequest("on_confirm"))

	port := "8080"
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleRequest(endpoint string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received request at /ondc/%s\n", endpoint)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, `{"message":"%s received"}`, endpoint)
	}
}
