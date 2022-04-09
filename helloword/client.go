package helloword

import (
	"context"
	"fmt"
	"github.com/grpc-example/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

const SerName = "psp-scale"

var conn *grpc.ClientConn

type Client struct {
	gclent GreeterClient
}

func NewClient() *Client {
	if conn == nil {
		initConn()
	}
	c := NewGreeterClient(conn)
	return &Client{
		gclent: c,
	}
}

func initConn() {
	var err error
	conn, err = grpc.Dial(
		fmt.Sprintf("%s:///%s", etcd.Resover.Scheme(), SerName),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

}

func (c *Client) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return c.gclent.SayHello(ctx, in)
}
