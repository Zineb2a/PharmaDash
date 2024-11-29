package main

import (
	"log"
	server "pharmaDashServer/router"

	_ "github.com/lib/pq"
)

func main() {

	// Server
	server, err := server.GetServerSingleton()
	if err != nil {
		log.Fatal("Could not start the server")
	}

	server.Router.Run(":3000")
}
