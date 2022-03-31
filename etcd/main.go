package main

//
//import (
//	"context"
//	"fmt"
//	clientv3 "go.etcd.io/etcd/client/v3"
//	"time"
//)
//
//func main() {
//	config := clientv3.Config{
//		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
//		DialTimeout: 5 * time.Second,
//	}
//	client, err := clientv3.New(config)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	timeoutCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
//	defer cancel()
//	_, err = client.Status(timeoutCtx, config.Endpoints[0])
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	defer client.Close()
//
//	kv := clientv3.NewKV(client)
//	var putReps *clientv3.PutResponse
//	var getReps *clientv3.GetResponse
//	//添加任务
//	if putReps, err = kv.Put(context.TODO(), "/cron/jobs/job1", "hello1"); err != nil {
//		panic(err)
//	}
//
//	if putReps, err = kv.Put(context.TODO(), "/cron/jobs/job2", "hello2"); err != nil {
//		panic(err)
//	}
//
//	fmt.Println(putReps.Header.Revision)
//
//	if getReps, err = kv.Get(context.TODO(), "/cron/jobs/job1"); err != nil {
//		panic(err)
//	}
//	fmt.Println(getReps)
//}
