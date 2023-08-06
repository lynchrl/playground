package bst

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSearch(t *testing.T) {
	root := testTree()

	type args struct {
		root *Node
		key  int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "key_exists",
			args: args{
				root: root,
				key:  7,
			},
			want: root.Right,
		},
		{
			name: "key_does_not_exist",
			args: args{
				root: root,
				key:  4,
			},
			want: nil,
		},
		{
			name: "nil_root",
			args: args{
				root: nil,
				key:  4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSearchI(t *testing.T) {
	root := testTree()

	type args struct {
		root *Node
		key  int
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "key_exists",
			args: args{
				root: root,
				key:  7,
			},
			want: root.Right,
		},
		{
			name: "key_does_not_exist",
			args: args{
				root: root,
				key:  4,
			},
			want: nil,
		},
		{
			name: "nil_root",
			args: args{
				root: nil,
				key:  4,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchI(tt.args.root, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	root := testTree()

	tests := []struct {
		name string
		root *Node
		want *Node
	}{
		{
			name: "nil_root",
			root: nil,
			want: nil,
		},
		{
			name: "valid_min",
			root: root,
			want: root.Left.Left,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMax(t *testing.T) {
	root := testTree()

	tests := []struct {
		name string
		root *Node
		want *Node
	}{
		{
			name: "nil_root",
			root: nil,
			want: nil,
		},
		{
			name: "valid_max",
			root: root,
			want: root.Right.Right,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuccessor(t *testing.T) {
	root := testTree()

	tests := []struct {
		name string
		node *Node
		want *Node
	}{
		{
			name: "nil_node",
			node: nil,
			want: nil,
		},
		{
			name: "root_successor",
			node: root,
			want: root.Right.Left,
		},
		{
			name: "leaf_successor",
			node: root.Right.Left,
			want: root.Right,
		},
		{
			name: "no_successor",
			node: root.Right.Right,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Successor(tt.node); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Successor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	want := &Node{Key: 5}
	want.Left = &Node{Key: 3, Parent: want}
	want.Right = &Node{Key: 7, Parent: want}

	want.Left.Left = &Node{Key: 2, Parent: want.Left}
	want.Left.Right = &Node{Key: 4, Parent: want.Left}

	got := Insert(nil, 5)
	Insert(got, 3)
	Insert(got, 7)
	Insert(got, 2)
	Insert(got, 4)

	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("Insert() mismatch (-want +got):\n%s", diff)
	}
}

func TestDelete(t *testing.T) {
	// TODO: Better tests/coverage.
	want := &Node{Key: 5}
	want.Left = &Node{Key: 3, Parent: want}
	want.Right = &Node{Key: 7, Parent: want}
	want.Left.Left = &Node{Key: 2, Parent: want.Left}
	want.Left.Right = &Node{Key: 4, Parent: want.Left}

	root := Insert(nil, 5)
	Insert(root, 3)
	Insert(root, 7)
	Insert(root, 2)
	Insert(root, 4)
	Insert(root, 6)

	Delete(root, root.Right.Left)

	if diff := cmp.Diff(root, want); diff != "" {
		t.Errorf("Delete() mismatch (-want +got):\n%s", diff)
	}
}

func BenchmarkSearch(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	// Prepare a BST with 1000 nodes
	root := generateBST(r, 1000)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Search for a random key in the BST
		key := r.Intn(1000)
		Search(root, key)
	}
}

func BenchmarkSearchI(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	// Prepare a BST with 1000 nodes
	root := generateBST(r, 1000)

	// Run the benchmark
	for i := 0; i < b.N; i++ {
		// Search for a random key in the BST
		key := r.Intn(1000)
		SearchI(root, key)
	}
}

// Helper function to generate a BST with n nodes
func generateBST(r *rand.Rand, n int) *Node {
	if n == 0 {
		return nil
	}

	// Generate a random key for the root node.
	root := &Node{Key: r.Intn(1000)}

	// Generate n-1 random keys for the left and right subtrees.
	for i := 0; i < n-1; i++ {
		key := r.Intn(1000)
		Insert(root, key)
	}

	return root
}

func testTree() *Node {
	root := Insert(nil, 5)
	Insert(root, 3)
	Insert(root, 7)
	Insert(root, 2)
	Insert(root, 6)
	Insert(root, 8)
	return root
}
