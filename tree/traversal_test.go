package tree_test

import (
	"reflect"
	"testing"

	"github.com/catninpo/dsa/tree"
)

func TestPreOrderSearch(t *testing.T) {
	bTree := newBinaryTree()

	values := tree.PreOrderTraversal(bTree)
	expectedOrder := []int{5, 8, 12, 16, 10}

	if !reflect.DeepEqual(values, expectedOrder) {
		t.Fatalf("[\u2717] return values not expected: got=%d want=%d",
			values, expectedOrder)
	}
	t.Logf("[\u2713] expected values in correct order")
}

func TestInOrderSearch(t *testing.T) {
	bTree := newBinaryTree()

	values := tree.InOrderTraversal(bTree)
	expectedOrder := []int{12, 8, 16, 5, 10}

	if !reflect.DeepEqual(values, expectedOrder) {
		t.Fatalf("[\u2717] return values not expected: got=%d want=%d",
			values, expectedOrder)
	}
	t.Logf("[\u2713] expected values in correct order")
}

func TestPostOrderSearch(t *testing.T) {
	bTree := newBinaryTree()

	values := tree.PostOrderTraversal(bTree)
	expectedOrder := []int{12, 16, 8, 10, 5}

	if !reflect.DeepEqual(values, expectedOrder) {
		t.Fatalf("[\u2717] return values not expected: got=%d want=%d",
			values, expectedOrder)
	}
	t.Logf("[\u2713] expected values in correct order")
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
