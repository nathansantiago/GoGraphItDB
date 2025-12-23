package main

import "slices"

type Graph struct {
	ID            string
	Name          string
	Nodes         map[string]*Node
	Relationships []*Relationship
}

// Nodes
func (g *Graph) AddNode(n *Node) {
	g.Nodes[n.ID] = n
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
	if idx != -1 {
		g.Relationships = append(g.Relationships[:idx], g.Relationships[idx+1:]...)
	}
}

// Node Degree
func (g *Graph) GetNodeDegree(nodeID string) int {
	count := 0
	for _, rel := range g.Relationships {
		if rel.StartNode.ID == nodeID || rel.EndNode.ID == nodeID {
			count++
		}
	}
	return count
}

// Traversing
func (g *Graph) GetNodeRelationships(nodeID string) []*Relationship {
	var relationships []*Relationship
	for _, rel := range g.Relationships {
		if rel.StartNode.ID == nodeID || rel.EndNode.ID == nodeID {
			relationships = append(relationships, rel)
		}
	}
	return relationships
}
