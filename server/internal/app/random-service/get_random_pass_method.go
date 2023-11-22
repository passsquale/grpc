package random_service

import (
	"context"
	random_service "github.com/passsquale/grpc/server/pkg/random-service"
	"math/rand"
)

func getString(length int64) string {
	temp := ""
	for i := int64(0); i < length; i++ {
		temp += string(byte(random(33, 127)))
	}
	return temp
}

func (*Implementation) GetRandomPass(ctx context.Context,
	r *random_service.RequestPass) (*random_service.RandomPass, error) {
	rand.New(rand.NewSource(r.GetSeed()))
	temp := getString(r.GetLength())
	return &random_service.RandomPass{
		Password: temp,
	}, nil
}
