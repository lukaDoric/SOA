package main

import (
	"context"
	"example/grpc/proto/product"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	productService := product.NewProductServiceClient(conn)

	getResp, err := productService.GetProduct(context.Background(), &product.GetProductRequest{Id: 1})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Product)
	}

	_, err = productService.DeleteProduct(context.Background(), &product.DeleteProductRequest{Id: 1})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("product deleted")
	}

	getResp, err = productService.GetProduct(context.Background(), &product.GetProductRequest{Id: 1})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResp.Product)
	}
}
