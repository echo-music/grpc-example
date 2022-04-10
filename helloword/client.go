package helloword

import (
	"context"
	"fmt"
	"github.com/grpc-example/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

const ServiceName = "psp-scale"

var client *Client

type Client struct {
	cli GreeterClient
}

// NewClient 创建grpc客户端
func NewClient() *Client {
	if client == nil {
		r := etcd.NewServiceDiscovery()
		resolver.Register(r)
		conn, err := grpc.Dial(
			fmt.Sprintf("%s:///%s", r.Scheme(), ServiceName),
			grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			panic(err)
		}
		c := NewGreeterClient(conn)
		client = &Client{
			cli: c,
		}
	}

	return client
}

// SayHello 调用方法
func (c *Client) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	return c.cli.SayHello(ctx, in)
}
