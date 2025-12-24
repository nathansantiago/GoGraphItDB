package main

import "slices"

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

func (n *Node) RemoveLabel(label string) {
	idx := slices.Index(n.Labels, label)
	n.Labels = append(n.Labels[:idx], n.Labels[idx+1:]...)
}

func (n *Node) SetProperty(key string, value interface{}) {
	n.Properties[key] = value
}

func (n *Node) GetProperty(key string) (interface{}, bool) {
	value, exists := n.Properties[key]
	return value, exists
}

func (n *Node) RemoveProperty(key string) {
	delete(n.Properties, key)
}

func (n *Node) GetLabels() []string {
	return n.Labels
}

func (n *Node) HasLabel(label string) bool {
	return slices.Contains(n.Labels, label)
}
