package registry

// 服务
type Service struct {
	Name  string
	Nodes []*Node
}

// 服务节点
type Node struct {
	Id   string
	Ip   string
	Port int
}
