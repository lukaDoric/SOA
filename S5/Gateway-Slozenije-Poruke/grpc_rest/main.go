package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ps "github.com/milossimic/grpc_rest/poststore"
	helloworldpb "github.com/milossimic/grpc_rest/proto/helloworld"
	tracer "github.com/milossimic/grpc_rest/tracer"
	otgo "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"io"
)

var grpcGatewayTag = otgo.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

func tracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := otgo.GlobalTracer().Extract(
			otgo.HTTPHeaders,
			otgo.HTTPHeadersCarrier(r.Header))
		if err == nil || err == otgo.ErrSpanContextNotFound {
			serverSpan := otgo.GlobalTracer().StartSpan(
				"ServeHTTP",
				// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(otgo.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}

type server struct {
	helloworldpb.UnimplementedGreeterServer
	store  *ps.PostStore
	tracer otgo.Tracer
	closer io.Closer
}

const name = "post_service"

func NewServer() (*server, error) {
	store, err := ps.New()
	if err != nil {
		return nil, err
	}

	tracer, closer := tracer.Init(name)
	otgo.SetGlobalTracer(tracer)
	return &server{
		store:  store,
		tracer: tracer,
		closer: closer,
	}, nil
}

func (s *server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *server) GetCloser() io.Closer {
	return s.closer
}

func (s *server) PostRequest(ctx context.Context, in *helloworldpb.CreatePostRequest) (*helloworldpb.Post, error) {
	return s.store.Post(ctx, in)
}

func (s *server) GetRequest(ctx context.Context, in *helloworldpb.GetPostRequest) (*helloworldpb.Post, error) {
	return s.store.Get(ctx, in.Post)
}

func (s *server) GetAllRequest(ctx context.Context, in *helloworldpb.EmptyRequest) (*helloworldpb.GetAllPosts, error) {
	return s.store.GetAll(ctx)
}

func (s *server) DeleteRequest(ctx context.Context, in *helloworldpb.DeletePostRequest) (*helloworldpb.Post, error) {
	return s.store.Delete(ctx, in.Post)
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()

	service, err := NewServer()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// Attach the Greeter service to the server
	helloworldpb.RegisterGreeterServer(s, service)
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		),
		grpc.WithStreamInterceptor(
			grpc_opentracing.StreamClientInterceptor(
				grpc_opentracing.WithTracer(otgo.GlobalTracer()),
			),
		),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = helloworldpb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: tracingWrapper(gwmux),
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
