/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"github.com/grpc-example/etcd"
	"github.com/grpc-example/protos"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	// Address 监听地址
	Address string = "localhost:8001"
	// Network 网络通信协议
	Network string = "tcp"
	// SerName 服务名称
	SerName string = "simple_grpc"
)

// EtcdEndpoints etcd地址
var EtcdEndpoints = []string{"localhost:2379"}

// server is used to implement helloworld.GreeterServer.
type server struct {
	protos.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *protos.HelloRequest) (*protos.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &protos.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protos.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())

	//把服务注册到etcd
	ser, err := etcd.NewServiceRegister(EtcdEndpoints, SerName, Address, 5)
	if err != nil {
		log.Fatalf("register service err: %v", err)
	}
	defer ser.Close()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}