package registry

import "time"

type Options struct {
	Addrs   []string
	TimeOut time.Duration
}

type Option func(opts *Options)

// WithTimeOut 设置超时
func WithTimeOut(timeout time.Duration) Option {
	return func(opts *Options) {
		opts.TimeOut = timeout
	}
}

// WithAddrs 设置服务地址
func WithAddrs(addrs []string) Option {
	return func(opts *Options) {
		opts.Addrs = addrs
	}
}
