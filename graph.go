package main

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

// Node Properties

func GetNodeProperties

func UpdateNodeProperties

func DeleteNodeProperties

func UpdateNodeProperty

func DeleteNodeProperty

// Relationships

// Relationship Properties

// Node Degree

// Traversing

// 