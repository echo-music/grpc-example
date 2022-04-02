package main

import (
	"context"
	"fmt"
	"github.com/grpc-example/etcd"
	"github.com/grpc-example/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

var (
	// EtcdEndpoints etcd地址
	EtcdEndpoints = []string{"localhost:2379"}
	// SerName 服务名称
	SerName = "simple_grpc"
)

func main() {

	r := etcd.NewServiceDiscovery(EtcdEndpoints)
	resolver.Register(r)
	// Set up a connection to the server.
	// 连接服务器
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", r.Scheme(), SerName),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()
	c := protos.NewGreeterClient(conn)

	// Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	//defer cancel()

	for i := 0; i < 1000; i++ {
		hello, err := c.SayHello(context.Background(), &protos.HelloRequest{Name: fmt.Sprintf("hello_%d", i)})
		if err != nil {
			return
		}
		fmt.Println(hello)
		time.Sleep(1 * time.Second)
	}

	select {}

}
