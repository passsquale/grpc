package random_service

import (
	"context"
	random_service "github.com/passsquale/grpc/server/pkg/random-service"
	"math/rand"
	"time"
)

func random(min, max int64) int64 {
	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)
	return rng.Int63n(max) + min
}

func (*Implementation) GetRandom(ctx context.Context, r *random_service.RandomParams) (*random_service.RandomInt, error) {
	return &random_service.RandomInt{
		Value: random(r.GetMin(), r.GetMax()),
	}, nil
}
