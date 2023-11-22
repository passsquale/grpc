package server

import (
	"context"
	"fmt"
	api "github.com/passsquale/grpc/server/internal/app/random-service"
	"github.com/passsquale/grpc/server/internal/config"
	random_service "github.com/passsquale/grpc/server/pkg/random-service"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type GrpcServer struct {
}

func NewGrpcServer() *GrpcServer {
	return &GrpcServer{}
}

func (s *GrpcServer) Start(cfg *config.Config) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcAddr := fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}
	defer l.Close()
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: time.Duration(cfg.MaxConnectionIdle) * time.Minute,
			Timeout:           time.Duration(cfg.Timeout) * time.Second,
			MaxConnectionAge:  time.Duration(cfg.MaxConnectionAge) * time.Minute,
			Time:              time.Duration(cfg.Timeout) * time.Minute,
		}))
	random_service.RegisterRandomServer(grpcServer, api.NewRandomService())
	go func() {
		log.Info().Msgf("GRPC Server is listening on: %s", grpcAddr)
		if err := grpcServer.Serve(l); err != nil {
			log.Fatal().Err(err).Msg("Failed running gRPC server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Info().Msgf("signal.Notify: %v", v)
	case done := <-ctx.Done():
		log.Info().Msgf("ctx.Done: %v", done)
	}

	grpcServer.GracefulStop()
	log.Info().Msgf("grpcServer shut down correctly")
	return nil
}
