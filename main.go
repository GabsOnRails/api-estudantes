package main

import (
	"log"

	"github.com/GabsOnRails/api-estudantes/api"
)

func main() {
	server := api.NewServer()
	server.ConfigureRoutes()
	if err := server.StartServer(); err != nil {
		log.Fatal(err)
	}
}
