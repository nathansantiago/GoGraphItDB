package main

import (
	"slices"
	"sync"
)

type Graph struct {
	ID            string
	Name          string
	Nodes         map[string]*Node
	Relationships []*Relationship
	mu            sync.RWMutex
}

func NewGraph(id string, name string) *Graph {
	return &Graph{
		ID:            id,
		Name:          name,
		Nodes:         make(map[string]*Node),
		Relationships: make([]*Relationship, 0),
	}
}

// Nodes
func (g *Graph) AddNode(n *Node) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Nodes[n.ID] = n
}

func (g *Graph) RemoveNodeByID(k string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.Nodes, k)
}

func (g *Graph) GetNodeByID(k string) *Node {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.Nodes[k]
}

// Relationships
func (g *Graph) AddRelationship(r *Relationship) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Relationships = append(g.Relationships, r)
}

func (g *Graph) RemoveRelationship(r *Relationship) {
	g.mu.Lock()
	defer g.mu.Unlock()
	idx := slices.Index(g.Relationships, r)
	if idx != -1 {
		g.Relationships = append(g.Relationships[:idx], g.Relationships[idx+1:]...)
	}
}

// Node Degree
func (g *Graph) GetNodeDegree(nodeID string) int {
	g.mu.RLock()
	defer g.mu.RUnlock()
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
	g.mu.RLock()
	defer g.mu.RUnlock()
	var relationships []*Relationship
	for _, rel := range g.Relationships {
		if rel.StartNode.ID == nodeID || rel.EndNode.ID == nodeID {
			relationships = append(relationships, rel)
		}
	}
	return relationships
}
