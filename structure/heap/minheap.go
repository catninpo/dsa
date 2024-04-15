package heap

import "errors"

type MinHeap[T string | int] struct {
	Data   []T
	Length int
}

func NewMinHeap[T int | string]() MinHeap[T] {
	return MinHeap[T]{
		Length: 0,
		Data:   make([]T, 0),
	}
}

func (mh *MinHeap[T]) Insert(value T) {
	mh.Data[mh.Length] = value
	mh.HeapifyUp(mh.Length)

	mh.Length++
}

func (mh *MinHeap[T]) Delete() (T, error) {
	if mh.Length == 0 {
		return *new(T), errors.New("heap is zero-length")
	}

	value := mh.Data[0]
	mh.Length--

	if mh.Length == 0 {
		mh.Data = make([]T, 0)
		return value, nil
	}

	mh.Data[0] = mh.Data[mh.Length]
	mh.HeapifyDown(0)

	return value, nil
}

func (mh *MinHeap[T]) HeapifyDown(index int) {
	if index >= mh.Length {
		return
	}

	lIndex := mh.LeftChild(index)

	if lIndex >= mh.Length {
		return
	}

	rIndex := mh.RightChild(index)

	lValue := mh.Data[lIndex]
	rValue := mh.Data[rIndex]
	value := mh.Data[index]

	switch {
	case lValue > rValue && value > rValue:
		mh.Data[index], mh.Data[rIndex] = rValue, value
		mh.HeapifyDown(rIndex)
	case rValue > lValue && value > lValue:
		mh.Data[index], mh.Data[lIndex] = lValue, value
		mh.HeapifyDown(lIndex)
	}
}

func (mh *MinHeap[T]) HeapifyUp(index int) {
	if index == 0 {
		return
	}

	parent := mh.Parent(index)
	parentValue := mh.Data[parent]

	value := mh.Data[index]

	if parentValue > value {
		mh.Data[index], mh.Data[parent] = parentValue, value
		mh.HeapifyUp(parent)
	}
}

func (mh *MinHeap[T]) Parent(index int) int {
	return (index - 1) / 2
}

func (mh *MinHeap[T]) LeftChild(index int) int {
	return (index * 2) + 1
}

func (mh *MinHeap[T]) RightChild(index int) int {
	return (index * 2) + 2
}
