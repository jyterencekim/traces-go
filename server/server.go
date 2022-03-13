package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jyterencekim/traces-go/server/handler"
)

func Run(port int) {
	addrPort := fmt.Sprintf(":%d", port)
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.DefaultHandler)

	log.Printf("Starting traces server at port %v", addrPort)
	if err := http.ListenAndServe(addrPort, mux); err != nil {
		log.Fatal(err)
	}
}
