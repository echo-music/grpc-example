package etcd

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
	"log"
	"sync"
)

const schema = "grpclb"

//ServiceDiscovery 服务发现
type ServiceDiscovery struct {
	cli        *clientv3.Client //etcd client
	cc         resolver.ClientConn
	serverList map[string]resolver.Address //服务列表
	lock       sync.Mutex
}

var serverList = make(map[string]resolver.Address) //服务列表

//NewServiceDiscovery  新建发现服务
func NewServiceDiscovery() resolver.Builder {
	return &ServiceDiscovery{cli: cli}
}

//Build 为给定目标创建一个新的`resolver`，当调用`grpc.Dial()`时执行
func (s *ServiceDiscovery) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	log.Println("Build")
	s.cc = cc
	s.serverList = make(map[string]resolver.Address)
	prefix := "/" + target.URL.Scheme + target.URL.Path + "/"
	//根据前缀获取现有的key
	resp, err := s.cli.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	fmt.Println("从etcd拉取所有数据")

	for _, ev := range resp.Kvs {
		s.SetServiceList(string(ev.Key), string(ev.Value))
	}
	state := resolver.State{
		Addresses:     s.getServices(),
		ServiceConfig: nil,
		Attributes:    nil,
	}
	if err = s.cc.UpdateState(state); err != nil {
		return nil, err
	}

	//监视前缀，修改变更的server
	go s.watcher(prefix)
	return s, nil
}

// ResolveNow 监视目标更新
func (s *ServiceDiscovery) ResolveNow(rn resolver.ResolveNowOptions) {
	log.Println("ResolveNow")
}

//Scheme return schema
func (s *ServiceDiscovery) Scheme() string {
	return schema
}

//Close 关闭
func (s *ServiceDiscovery) Close() {
	log.Println("Close")
	s.cli.Close()
}

//watcher 监听前缀
func (s *ServiceDiscovery) watcher(prefix string) {
	rch := s.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	log.Printf("watching prefix:%s now...", prefix)
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case mvccpb.PUT: //新增或修改
				s.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE: //删除
				s.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

//SetServiceList 新增服务地址
func (s *ServiceDiscovery) SetServiceList(key, val string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	fmt.Println(key, "key")
	s.serverList[key] = resolver.Address{Addr: val}
	state := resolver.State{
		Addresses:     s.getServices(),
		ServiceConfig: nil,
		Attributes:    nil,
	}
	s.cc.UpdateState(state)
	log.Println("put key :", key, "val:", val)
}

//DelServiceList 删除服务地址
func (s *ServiceDiscovery) DelServiceList(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.serverList, key)
	state := resolver.State{
		Addresses:     s.getServices(),
		ServiceConfig: nil,
		Attributes:    nil,
	}
	s.cc.UpdateState(state)

	log.Println("del key:", key)
}

//GetServices 获取服务地址
func (s *ServiceDiscovery) getServices() []resolver.Address {

	addrs := make([]resolver.Address, 0, len(s.serverList))
	for _, v := range s.serverList {
		addrs = append(addrs, v)
	}

	return addrs
}
