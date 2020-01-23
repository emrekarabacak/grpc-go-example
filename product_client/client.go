package main

import (
	"context"
	"log"

	"github.com/emrekarabacak/grpc-go-example/product"
	"google.golang.org/grpc"
)

func main() {

	clientConn, err := grpc.Dial("0.0.0.0:55001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial %v", err)
	}
	defer clientConn.Close()

	client := product.NewProductServiceClient(clientConn)

	input := &product.ProductRequest{
		ProductId: "product-1",
	}
	response, err := client.GetDetails(context.Background(), input)
	if err != nil {
		log.Fatalf("Request failed with %v", err)
	}

	log.Printf("Request completed. Response is %v", response)

}
