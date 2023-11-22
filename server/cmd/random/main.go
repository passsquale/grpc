package main

import (
	"flag"
	"github.com/passsquale/grpc/server/internal/config"
	"github.com/passsquale/grpc/server/internal/server"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := config.ReadConfigYML("config.yml"); err != nil {
		log.Fatal().Err(err).Msg("Failed init configuration")
	}
	cfg := config.GetConfigInstance()
	flag.Parse()
	if err := server.NewGrpcServer().Start(&cfg); err != nil {
		log.Error().Err(err).Msg("Failed creating gRPC server")
		return
	}
}
