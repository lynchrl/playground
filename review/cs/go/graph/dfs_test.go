// Package graph includes types and functions for working with simple graphs.
package graph

import (
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/google/go-cmp/cmp"
)

func TestUndirected_AddEdge(t *testing.T) {
	type edge struct {
		from int
		to   int
	}
	tests := []struct {
		name    string
		edges   []edge
		wantAdj map[int]mapset.Set[int]
	}{
		{
			name: "one edge",
			edges: []edge{
				{from: 0, to: 1},
			},
			wantAdj: map[int]mapset.Set[int]{
				0: mapset.NewSet[int](1),
				1: mapset.NewSet[int](0)},
		},
		{
			name: "multiple edges",
			edges: []edge{
				{from: 0, to: 1},
				{from: 1, to: 2},
				{from: 0, to: 4},
			},
			wantAdj: map[int]mapset.Set[int]{
				0: mapset.NewSet[int](1, 4),
				1: mapset.NewSet[int](0, 2),
				2: mapset.NewSet[int](1),
				4: mapset.NewSet[int](0)},
		},
		{
			name: "disjoint",
			edges: []edge{
				{from: 0, to: 1},
				{from: 1, to: 2},
				{from: 0, to: 4},
				{from: 7, to: 8},
			},
			wantAdj: map[int]mapset.Set[int]{
				0: mapset.NewSet[int](1, 4),
				1: mapset.NewSet[int](0, 2),
				2: mapset.NewSet[int](1),
				4: mapset.NewSet[int](0),
				7: mapset.NewSet[int](8),
				8: mapset.NewSet[int](7),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUndirected()
			for _, e := range tt.edges {
				u.AddEdge(e.from, e.to)
			}
			if diff := cmp.Diff(tt.wantAdj, u.adj); diff != "" {
				t.Errorf("Undirected.AddEdge() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}

func TestUndirected_GetNeighbors(t *testing.T) {
	type edge struct {
		from int
		to   int
	}
	tests := []struct {
		name  string
		edges []edge
		node  int
		want  []int
	}{
		{
			name: "some edges",
			edges: []edge{
				{from: 0, to: 1},
				{from: 1, to: 2},
				{from: 0, to: 4},
				{from: 7, to: 8},
			},
			node: 0,
			want: []int{1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewUndirected()
			for _, e := range tt.edges {
				u.AddEdge(e.from, e.to)
			}
			got := u.GetNeighbors(tt.node)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Undirected.GetNeighbors() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
