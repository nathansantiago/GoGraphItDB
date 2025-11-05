package main

import "slices"

type Graph struct {
	ID string
	Name string
	Nodes map[string]*Node
	Relationships []*Relationship
}

// Nodes
func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) RemoveNodeByID(k string) {
	delete(g.Nodes, k)
}

func (g *Graph) GetNodeByID(k string) *Node {
	return g.Nodes[k]
}

// Relationships
func (g *Graph) AddRelationship(r *Relationship) {
	g.Relationships = append(g.Relationships, r)
}

func (g *Graph) RemoveRelationship(r *Relationship) {
	idx := slices.Index(g.Relationships, r)
	g.Relationships = 
}

// Node Degree

// Traversing

// 