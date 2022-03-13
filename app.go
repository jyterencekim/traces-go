package main

import (
	"github.com/jyterencekim/traces-go/server"
)

func main() {
	// TODO(@jyterencekim): make the port configurable
	const PORT = 8080

	server.Run(PORT)
}
