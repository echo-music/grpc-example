package main

import (
	"github.com/grpc-example/helloword"
)

var (
	// SerName 服务名称
	SerName = "psp-scale"
)

func main() {

	helloword.SayHello(0)

	select {}

}

func doWork(i int) {

}
