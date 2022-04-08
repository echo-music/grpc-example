package main

import (
	"github.com/grpc-example/helloword"
	"time"
)

var (
	// SerName 服务名称
	SerName = "psp-scale"
)

func main() {

	for i := 0; i < 10; i++ {
		helloword.SayHello(i)
		time.Sleep(1 * time.Second)

	}
	select {}

}

func doWork(i int) {

}
