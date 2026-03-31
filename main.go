package main

import (
	"github.com/rs/zerolog/log"

	"github.com/GabsOnRails/api-estudantes/api"
)

func main() {
	server := api.NewServer()
	server.ConfigureRoutes()
	if err := server.StartServer(); err != nil {
		log.Fatal().Err(err).Msgf("Failed to start server: %v", err.Error())
	}
}
