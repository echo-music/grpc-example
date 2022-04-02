package registry

import (
	"context"
	"fmt"
	"sync"
)

// 插件管理初始化
var plugMgr = &PluginMgr{
	plugins: make(map[string]Registry),
}

// 插件管理
type PluginMgr struct {
	plugins map[string]Registry
	lock    sync.Mutex
}

// registerPlugin 注册插件
func (p *PluginMgr) registerPlugin(plugin Registry) (err error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	if _, ok := p.plugins[plugin.Name()]; ok {
		err = fmt.Errorf("duplicate registry plugin")
		return
	}
	p.plugins[plugin.Name()] = plugin
	return
}

// initRegister 初始化注册中心
func (p *PluginMgr) initRegister(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	registry, ok := p.plugins[name]
	if !ok {
		err = fmt.Errorf("plugin %s not exists", name)
		return
	}
	err = registry.Init(ctx, opts...)
	return
}

// RegisterPlugin 注册插件
func RegisterPlugin(plugin Registry) (err error) {
	return plugMgr.registerPlugin(plugin)
}

// InitRegister 初始化注册中心
func InitRegister(ctx context.Context, name string, opts ...Option) (registry Registry, err error) {
	return plugMgr.initRegister(ctx, name, opts...)
}
