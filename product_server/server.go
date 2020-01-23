package main

import (
	"context"
	"log"
	"net"

	"github.com/emrekarabacak/grpc-go-example/product"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetDetails(ctx context.Context, req *product.ProductRequest) (*product.ProductResponse, error) {
	return &product.ProductResponse{}, nil
}

func main() {
	log.Printf("Starting grpc server...")
	lis, err := net.Listen("tcp", "0.0.0.0:55001")
	if err != nil {
		log.Fatalf("Failed to listen :  %v", err)
	}

	grpcServer := grpc.NewServer()
	product.RegisterProductServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cannot start grpc server %v", err)
	}
}
