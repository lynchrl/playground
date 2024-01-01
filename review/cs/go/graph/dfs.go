// Package graph includes types and functions for working with simple graphs.
package graph

import (
	mapset "github.com/deckarep/golang-set/v2"
)

// Undirected is an undirected graph with integer nodes. Create with NewUndirected.
type Undirected struct {
	adj map[int]mapset.Set[int]
}

// NewUndirected creates a new undirected graph.
func NewUndirected() *Undirected {
	return &Undirected{adj: map[int]mapset.Set[int]{}}
}

// AddEdge adds an edge between from and to.
func (u *Undirected) AddEdge(from, to int) {
	if u.adj[from] == nil {
		u.adj[from] = mapset.NewSet[int]()
	}
	if u.adj[to] == nil {
		u.adj[to] = mapset.NewSet[int]()
	}
	u.adj[from].Add(to)
	u.adj[to].Add(from)
}

// GetNeighbors returns the neighbors of the given node.
func (u *Undirected) GetNeighbors(node int) []int {
	return u.adj[node].ToSlice()
}
