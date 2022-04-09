package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

var (
	cli           *clientv3.Client //etcd client
	EtcdEndpoints = []string{"localhost:2379"}
)

// init  初始化etcd客户端
func init() {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   EtcdEndpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = cli.Status(timeoutCtx, EtcdEndpoints[0])
	if err != nil {
		panic(err)
	}

}
