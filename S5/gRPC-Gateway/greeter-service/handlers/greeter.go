package handlers

import (
	"context"
	"example/grpc/proto/greeter"
	"fmt"
)

type GreeterHandler struct {
	greeter.UnimplementedGreeterServiceServer
}

func (h GreeterHandler) Greet(ctx context.Context, request *greeter.Request) (*greeter.Response, error) {
	return &greeter.Response{
		Greeting: fmt.Sprintf("Hi %s!", request.Name),
	}, nil
}
