// Package bst implements a simple binary search tree.
package bst

// Node is a node in a BST.
type Node struct {
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// Search searches for a key in the BST.
// Returns the Node if found, or nil otherwise.
func Search(root *Node, key int) *Node {
	if root == nil || root.Key == key {
		return root
	}

	if key < root.Key {
		return Search(root.Left, key)
	} else {
		return Search(root.Right, key)
	}
}

// SearchI searches for a key in the BST using an iterator.
// Returns the Node if found, or nil otherwise.
func SearchI(root *Node, key int) *Node {
	for root != nil && root.Key != key {
		if key < root.Key {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}
