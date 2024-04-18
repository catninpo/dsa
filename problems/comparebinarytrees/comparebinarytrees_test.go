package comparebinarytrees_test

import (
	"testing"

	"github.com/catninpo/dsa/problems/comparebinarytrees"
	"github.com/catninpo/dsa/tree"
)

func TestCompare(t *testing.T) {
	a, b := newBinaryTree(), newBinaryTree()

	eq := comparebinarytrees.Compare(a, b)

	if eq != true {
		t.Fatalf("trees should match")
	}
}

func newBinaryTree() *tree.BinaryNode[int] {
	return &tree.BinaryNode[int]{
		Value: 5,
		Left: &tree.BinaryNode[int]{
			Value: 8,
			Left:  &tree.BinaryNode[int]{Value: 12},
			Right: &tree.BinaryNode[int]{Value: 16},
		},
		Right: &tree.BinaryNode[int]{Value: 10},
	}
}
