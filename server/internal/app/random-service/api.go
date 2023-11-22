package random_service

import desc "github.com/passsquale/grpc/server/pkg/random-service"

type Implementation struct {
	desc.UnimplementedRandomServer
}

func NewRandomService() desc.RandomServer {
	return &Implementation{}
}
