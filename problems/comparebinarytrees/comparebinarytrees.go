package comparebinarytrees

import (
	"github.com/catninpo/dsa/tree"
)

func Compare(a *tree.BinaryNode[int], b *tree.BinaryNode[int]) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.Value != b.Value {
		return false
	}

	return Compare(a.Left, b.Left) && Compare(a.Right, b.Right)
}
