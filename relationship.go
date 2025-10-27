package main

type Relationship struct {
	ID         string
	Type       string
	StartNode  *Node
	EndNode    *Node
	Properties map[string]interface{}
}

func NewRelationship(id string, relType string, startNode *Node, endNode *Node) *Relationship {
	return &Relationship{
		ID:         id,
		Type:       relType,
		StartNode:  startNode,
		EndNode:    endNode,
		Properties: make(map[string]interface{}),
	}
}

func (r *Relationship) SetProperty(key string, value interface{}) {
	r.Properties[key] = value
}

func (r *Relationship) GetProperty(key string) (interface{}, bool) {
	value, exists := r.Properties[key]
	return value, exists
}
