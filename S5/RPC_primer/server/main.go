package main

import (
	"context"
	"example/grpc/proto/product"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	product.RegisterProductServiceServer(grpcServer, Server{products: products})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
}

type Server struct {
	product.UnimplementedProductServiceServer
	products map[int32]*product.Product
}

func (s Server) GetProduct(ctx context.Context, request *product.GetProductRequest) (*product.GetProductResponse, error) {
	p, ok := s.products[request.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	response := &product.GetProductResponse{
		Product: p,
	}
	return response, nil
}

func (s Server) UpsertProduct(ctx context.Context, request *product.UpsertProductRequest) (*product.UpsertProductResponse, error) {
	s.products[request.Product.Id] = request.Product
	response := &product.UpsertProductResponse{
		Product: s.products[request.Product.Id],
	}
	return response, nil
}

func (s Server) DeleteProduct(ctx context.Context, request *product.DeleteProductRequest) (*product.DeleteProductResponse, error) {
	if _, ok := s.products[request.Id]; !ok {
		return nil, status.Error(codes.NotFound, "product not found")
	}
	delete(s.products, request.Id)
	return &product.DeleteProductResponse{}, nil
}

var products = map[int32]*product.Product{
	1: {
		Id:          1,
		Category:    product.Product_CLOTHES,
		Description: "green sweatshirt",
		Price:       1999.99,
	},
	2: {
		Id:          2,
		Category:    product.Product_ELECTRONICS,
		Description: "tablet",
		Price:       12345.67,
	},
	3: {
		Id:          3,
		Category:    product.Product_BOOKS,
		Description: "Stephen King's The Shining",
		Price:       700,
	},
}
