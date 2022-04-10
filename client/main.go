package main

import (
	"context"
	"fmt"
	"github.com/grpc-example/helloword"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	client := helloword.NewClient()
	reply, _ := client.SayHello(context.Background(), &helloword.HelloRequest{
		Name:   "张三",
		Gender: "女",
	})

	pingReply, _ := client.Ping(context.Background(), &helloword.PingRequest{})

	msg := fmt.Sprintf("%s,%s", reply.Message, pingReply.Message)
	fmt.Fprintln(w, msg)
}

func main() {

	http.HandleFunc("/", IndexHandler)     //设置访问的路由
	err := http.ListenAndServe(":80", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
