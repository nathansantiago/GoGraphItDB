package main

type Node struct {
	ID         string
	Labels     []string
	Properties map[string]interface{}
}

func NewNode(id string) *Node {
	return &Node{
		ID:         id,
		Labels:     make([]string, 0),
		Properties: make(map[string]interface{}),
	}
}

func (n *Node) AddLabel(label string) {
	n.Labels = append(n.Labels, label)
}

func (n *Node) SetProperty(key string, value interface{}) {
	n.Properties[key] = value
}

func (n *Node) GetProperty(key string) (interface{}, bool) {
	value, exists := n.Properties[key]
	return value, exists
}
