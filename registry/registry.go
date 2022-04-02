package registry

import "context"

// Registry 插件抽象
type Registry interface {
	Name() string
	Init(ctx context.Context, opts ...Option) (err error)
	Register(ctx context.Context, service *Service) (err error)
	UnRegister(ctx context.Context, service *Service) (err error)
}
