package main

import (
	"context"
	"fmt"
	random_service "github.com/passsquale/grpc/server/pkg/random-service"
	"google.golang.org/grpc"
	"math/rand"
	"os"
	"time"
)

var port = ":8082"

func AskingDateTime(ctx context.Context, m random_service.RandomClient) (*random_service.DateTime, error) {
	request := &random_service.RequestDateTime{
		Value: "Please send me the date and time",
	}

	return m.GetDate(ctx, request)
}

func AskPass(ctx context.Context, m random_service.RandomClient, seed int64, length int64) (*random_service.RandomPass, error) {
	request := &random_service.RequestPass{
		Seed:   seed,
		Length: length,
	}

	return m.GetRandomPass(ctx, request)
}

func AskRandom(ctx context.Context, m random_service.RandomClient, min int64, max int64) (*random_service.RandomInt, error) {
	request := &random_service.RandomParams{
		Min: min,
		Max: max,
	}

	return m.GetRandom(ctx, request)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Using default port:", port)
	} else {
		port = os.Args[1]
	}

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial:", err)
		return
	}

	rand.Seed(time.Now().Unix())

	client := random_service.NewRandomClient(conn)
	r, err := AskingDateTime(context.Background(), client)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Server Date and Time:", r.Value)

	p, err := AskPass(context.Background(), client, 100, 10)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Password:", p.Password)

	i, err := AskRandom(context.Background(), client, 100, 100)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Integer 1:", i.Value)

	k, err := AskRandom(context.Background(), client, 100, 1000)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Random Integer 2:", k.Value)
}
