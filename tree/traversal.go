package tree

func PreOrderWalk(curr *BinaryNode[int], path []int) []int {
	if curr == nil {
		return path
	}

	path = append(path, curr.Value)

	path = PreOrderWalk(curr.Left, path)
	path = PreOrderWalk(curr.Right, path)

	return path
}

func PreOrderTraversal(head *BinaryNode[int]) []int {
	return PreOrderWalk(head, make([]int, 0))
}

func InOrderWalk(curr *BinaryNode[int], path []int) []int {
	if curr == nil {
		return path
	}

	path = InOrderWalk(curr.Left, path)
	path = append(path, curr.Value)
	path = InOrderWalk(curr.Right, path)

	return path
}

func InOrderTraversal(head *BinaryNode[int]) []int {
	return InOrderWalk(head, make([]int, 0))
}

func PostOrderWalk(curr *BinaryNode[int], path []int) []int {
	if curr == nil {
		return path
	}

	path = PostOrderWalk(curr.Left, path)
	path = PostOrderWalk(curr.Right, path)

	path = append(path, curr.Value)

	return path
}

func PostOrderTraversal(head *BinaryNode[int]) []int {
	return PostOrderWalk(head, make([]int, 0))
}
