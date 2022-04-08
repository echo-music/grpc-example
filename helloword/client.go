package helloword

import (
	"context"
	"fmt"
	"github.com/grpc-example/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

var SerName = "psp-scale"

func getConn() *grpc.ClientConn {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", etcd.Resover.Scheme(), SerName),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	return conn
	//defer conn.Close()
}

func SayHello(i int) {
	conn := getConn()

	c := NewGreeterClient(conn)
	//Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10000000)
	//defer cancel()
	defer conn.Close()
	for i := 0; i < 1000; i++ {
		hello, err := c.SayHello(context.Background(), &HelloRequest{Name: fmt.Sprintf("hello_%d", i)})
		if err != nil {
			return
		}
		fmt.Println(hello)
		time.Sleep(time.Second * 1)
	}

}
