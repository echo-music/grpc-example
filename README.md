# grpc-example
grpc,grpc-etcd,grpc-register,grpc-discovery

##说明
 支持服务注册和服务发现的小demo,仅供学习使用。
### 1、安装protobuf编辑器(mac下安装)
>  brew install protobuf
### 2、安装protobuf和grpc插件
> go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 
> go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 
### 3、编写proto文件
```
 // Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

syntax = "proto3";

// option go_package = "google.golang.org/grpc/examples/helloworld/helloworld";
option go_package = "grpc-example/helloword;helloword";


package protos;

// The greeting service definition.
service Greeter {
    // Sends a greeting
    rpc SayHello (HelloRequest) returns (HelloReply) {}
    rpc Ping (PingRequest) returns (PingReply) {}
}

// The request message containing the user's name.
message HelloRequest {
    string name = 1;
    string gender = 2;
}

// The response message containing the greetings
message HelloReply {
    string message = 1;
}

message PingRequest {

}

message PingReply{
    string message = 1;
}
```

### 4、生成grpc客户端和服务端代码
> protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
./helloword.proto
