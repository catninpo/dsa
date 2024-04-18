package tree

func BFS(head *BinaryNode[int], needle int) bool {
	queue := []*BinaryNode[int]{head}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Value == needle {
			return true
		}

		if curr.Left != nil {
			queue = append([]*BinaryNode[int]{curr.Left}, queue...)
		}

		if curr.Right != nil {
			queue = append([]*BinaryNode[int]{curr.Right}, queue...)
		}
	}

	return false
}
