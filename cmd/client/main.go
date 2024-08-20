package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/Lutefd/grpc-example/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	serverAddr := flag.String("server", "localhost:2706", "server address in format host:port")
	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := grpc.NewClient(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalculatorClient(conn)
	r, err := c.Add(ctx, &pb.CalculationRequest{A: 1, B: 2})
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	fmt.Printf("1 + 2 = %d\n", r.Result)

	r, err = c.Subtract(ctx, &pb.CalculationRequest{A: 1, B: 2})
	if err != nil {
		log.Fatalf("could not subtract: %v", err)
	}
	fmt.Printf("1 - 2 = %d\n", r.Result)

	r, err = c.Multiply(ctx, &pb.CalculationRequest{A: 1, B: 2})
	if err != nil {
		log.Fatalf("could not multiply: %v", err)
	}
	fmt.Printf("1 * 2 = %d\n", r.Result)

	r, err = c.Divide(ctx, &pb.CalculationRequest{A: 1, B: 2})
	if err != nil {
		log.Fatalf("could not divide: %v", err)
	}
	fmt.Printf("1 / 2 = %d\n", r.Result)

	r, err = c.Sum(ctx, &pb.NumbersRequest{
		Numbers: []int64{1, 2, 3, 4, 5},
	})
	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}
	fmt.Printf("1 + 2 + 3 + 4 + 5 = %d\n", r.Result)
}
