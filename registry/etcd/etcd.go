package etcd

import (
	"context"
	"github.com/grpc-example/registry"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdRegister struct {
	options   *registry.Options
	client    *clientv3.Client
	serviceCh chan *registry.Service
}

var (
	etcdRegistry EtcdRegister = EtcdRegister{
		serviceCh: make(chan *registry.Service, 8),
	}
)

func init() {
	err := registry.RegisterPlugin(&etcdRegistry)
	if err != nil {
		panic(err)
	}
}

func (e *EtcdRegister) Name() string {
	return "etcd"
}
func (e *EtcdRegister) Init(ctx context.Context, opts ...registry.Option) (err error) {
	return nil
}

func (e *EtcdRegister) Register(ctx context.Context, service *registry.Service) (err error) {
	return nil
}

func (e *EtcdRegister) UnRegister(ctx context.Context, service *registry.Service) (err error) {
	return nil
}
