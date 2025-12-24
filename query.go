package main

import (
	"reflect"
)

// FindNodesByLabel returns all nodes that have the given label.
func FindNodesByLabel(g *Graph, label string) []*Node {
	result := make([]*Node, 0)
	for _, n := range g.Nodes {
		if n.HasLabel(label) {
			result = append(result, n)
		}
	}
	return result
}

// FindNodesByProperties returns all nodes that match the given property key/value.
func FindNodesByProperties(g *Graph, key string, value interface{}) []*Node {
	result := make([]*Node, 0)
	for _, n := range g.Nodes {
		if v, ok := n.GetProperty(key); ok {
			if reflect.DeepEqual(v, value) {
				result = append(result, n)
			}
		}
	}
	return result
}

// FindRelationshipByType returns all relationships that match the given type.
func FindRelationshipByType(g *Graph, relationshipType string) []*Relationship {
	result := make([]*Relationship, 0)
	for _, r := range g.Relationships {
		if r.Type == relationshipType {
			result = append(result, r)
		}
	}
	return result
}

// FindPath returns a path (as a list of nodes) from startID to endID using BFS.
// If no path exists, returns nil.
func FindPath(g *Graph, startID string, endID string) []*Node {
	if startID == endID {
		if start := g.GetNodeByID(startID); start != nil {
			return []*Node{start}
		}
		return nil
	}

	if g.GetNodeByID(startID) == nil || g.GetNodeByID(endID) == nil {
		return nil
	}

	visited := make(map[string]bool)
	prev := make(map[string]string) // nodeID -> previous nodeID
	queue := []string{startID}
	visited[startID] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, rel := range g.GetNodeRelationships(current) {
			var neighbor *Node
			if rel.StartNode.ID == current {
				neighbor = rel.EndNode
			} else {
				neighbor = rel.StartNode
			}

			if neighbor == nil {
				continue
			}

			if !visited[neighbor.ID] {
				visited[neighbor.ID] = true
				prev[neighbor.ID] = current
				if neighbor.ID == endID {
					// reconstruct path
					return reconstructPath(g, startID, endID, prev)
				}
				queue = append(queue, neighbor.ID)
			}
		}
	}
	return nil
}

func reconstructPath(g *Graph, startID, endID string, prev map[string]string) []*Node {
	pathIDs := []string{}
	for at := endID; at != ""; at = prev[at] {
		pathIDs = append(pathIDs, at)
		if at == startID {
			break
		}
	}
	// reverse
	for i, j := 0, len(pathIDs)-1; i < j; i, j = i+1, j-1 {
		pathIDs[i], pathIDs[j] = pathIDs[j], pathIDs[i]
	}

	// convert to nodes
	path := make([]*Node, 0, len(pathIDs))
	for _, id := range pathIDs {
		if n := g.GetNodeByID(id); n != nil {
			path = append(path, n)
		} else {
			// if a node is missing unexpectedly, abort
			return nil
		}
	}
	return path
}
