package random_service

import (
	"context"
	random_service "github.com/passsquale/grpc/server/pkg/random-service"
	"time"
)

func (*Implementation) GetDate(ctx context.Context,
	r *random_service.RequestDateTime) (*random_service.DateTime, error) {
	currentTime := time.Now()
	return &random_service.DateTime{Value: currentTime.String()}, nil
}
