package main

import "github.com/grpc-example/helloword"

var (
	// SerName 服务名称
	SerName = "psp-scale"
)

func main() {

	for i := 0; i < 100; i++ {
		helloword.SayHello(i)
	}

	select {}

}

func doWork(i int) {

}
