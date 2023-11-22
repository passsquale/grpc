package random_service

import random_service "github.com/passsquale/grpc/server/pkg/random-service"

type Implementation struct {
	random_service.UnimplementedRandomServer
}

func NewRandomService() random_service.RandomServer {
	return &Implementation{}
}
