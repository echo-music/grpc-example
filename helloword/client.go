package helloword

import (
	"context"
	"fmt"
	"github.com/grpc-example/etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials/insecure"
)

var SerName = "psp-scale"
var Conn *grpc.ClientConn

func init() {
	var err error
	Conn, err = grpc.Dial(
		fmt.Sprintf("%s:///%s", etcd.Resover.Scheme(), SerName),
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

}

func SayHello(i int) {

	c := NewGreeterClient(Conn)
	//Contact the server and print out its response.
	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10000000)
	//defer cancel()

	hello, err := c.SayHello(context.Background(), &HelloRequest{Name: fmt.Sprintf("hello_%d", i)})
	if err != nil {
		return
	}
	fmt.Println(hello, "i=", i)
}
