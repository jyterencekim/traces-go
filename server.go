package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// TODO(@jyterencekim): make the port configurable
	const PORT = 8080
	addrPort := fmt.Sprintf(":%d", PORT)
	mux := http.NewServeMux()

	mux.HandleFunc("/", defaultHandler)

	log.Printf("Starting traces server at port %v", addrPort)
	if err := http.ListenAndServe(addrPort, mux); err != nil {
		log.Fatal(err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome: %v", r.URL.RawPath)
}
