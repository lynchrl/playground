package bst

import (
	"math/rand"
	"reflect"
	"testing"
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
		insert(root, key)
	}

	return root
}

// Helper function to insert a key into the BST.
func insert(root *Node, key int) *Node {
	if root == nil {
		return &Node{Key: key}
	}

	if key < root.Key {
		root.Left = insert(root.Left, key)
		root.Left.Parent = root
	} else {
		root.Right = insert(root.Right, key)
		root.Right.Parent = root
	}

	return root
}

func testTree() *Node {
	root := &Node{Key: 5}
	root.Left = &Node{Key: 3}
	root.Left.Left = &Node{Key: 2}
	root.Right = &Node{Key: 7}
	root.Right.Left = &Node{Key: 6}
	root.Right.Right = &Node{Key: 8}
	return root
}
