package poststore

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	helloworldpb "github.com/milossimic/grpc_rest/proto/helloworld"
	tracer "github.com/milossimic/grpc_rest/tracer"
	"os"
)

const (
	posts = "posts/%s"
	all   = "posts"
)

type PostStore struct {
	cli *api.Client
}

func New() (*PostStore, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &PostStore{
		cli: client,
	}, nil
}

func (ps *PostStore) Get(ctx context.Context, id string) (*helloworldpb.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "Get")
	defer span.Finish()

	kv := ps.cli.KV()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	pair, _, err := kv.Get(constructKey(ctx, id), nil)
	if err != nil {
		return nil, err
	}

	post := &helloworldpb.Post{}
	err = proto.Unmarshal(pair.Value, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *PostStore) GetAll(ctx context.Context) (*helloworldpb.GetAllPosts, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetAll")
	defer span.Finish()

	kv := ps.cli.KV()
	data, _, err := kv.List(all, nil)
	if err != nil {
		return nil, err
	}

	posts := []*helloworldpb.Post{}
	for _, pair := range data {
		post := &helloworldpb.Post{}
		err = proto.Unmarshal(pair.Value, post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return &helloworldpb.GetAllPosts{
		Posts: posts,
	}, nil
}

func (ps PostStore) Post(ctx context.Context, post *helloworldpb.CreatePostRequest) (*helloworldpb.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "Post")
	defer span.Finish()

	kv := ps.cli.KV()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	sid, rid := generateKey(ctx)
	post.Post.Id = rid

	data, err := proto.Marshal(post.Post)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return post.Post, nil
}

func (ps *PostStore) Delete(ctx context.Context, id string) (*helloworldpb.Post, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "Delete")
	defer span.Finish()

	kv := ps.cli.KV()

	ctx = tracer.ContextWithSpan(context.Background(), span)
	_, err := kv.Delete(constructKey(ctx, id), nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func generateKey(ctx context.Context) (string, string) {
	span := tracer.StartSpanFromContext(ctx, "generateKey")
	defer span.Finish()

	id := uuid.New().String()
	return fmt.Sprintf(posts, id), id
}

func constructKey(ctx context.Context, id string) string {
	span := tracer.StartSpanFromContext(ctx, "constructKey")
	defer span.Finish()

	return fmt.Sprintf(posts, id)
}
