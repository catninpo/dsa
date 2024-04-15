package doublelinkedlist

import "errors"

var (
	ErrItemNotFound  = errors.New("item not found")
	ErrIndexTooLarge = errors.New("index is larger than length of list")
)

type dllNode[T comparable] struct {
	Value T
	Prev  *dllNode[T]
	Next  *dllNode[T]
}

type DoubleLinkedList[T comparable] struct {
	Head   *dllNode[T]
	Tail   *dllNode[T]
	Length int
}

func New[T comparable]() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{
		Length: 0,
	}
}

func (dll *DoubleLinkedList[T]) Prepend(item T) {
	dll.Length++

	if dll.Head == nil {
		dll.Head = &dllNode[T]{
			Value: item,
		}
		dll.Tail = dll.Head
		return
	}

	node := dllNode[T]{
		Value: item,
		Next:  dll.Head,
	}

	dll.Head.Prev = &node
	dll.Head = &node
}

func (dll *DoubleLinkedList[T]) InsertAt(item T, index int) error {
	if index > dll.Length {
		return ErrIndexTooLarge
	}

	if index == dll.Length {
		dll.Append(item)
	}

	if index == 0 {
		dll.Prepend(item)
		return nil
	}

	dll.Length++

	curr := dll.Head
	for i := 0; curr != nil && i < index; i++ {
		curr = curr.Next
	}

	node := dllNode[T]{
		Value: item,
		Next:  curr,
		Prev:  curr.Prev,
	}

	if node.Prev != nil {
		curr.Prev.Next = &node
	}

	return nil
}

func (dll *DoubleLinkedList[T]) Append(item T) {
	dll.Length++
	node := dllNode[T]{Value: item}

	if dll.Tail == nil {
		dll.Head = &node
		dll.Tail = &node
		return
	}

	node.Prev = dll.Tail
	dll.Tail.Next = &node
	dll.Tail = &node
}

func (dll *DoubleLinkedList[T]) Remove(item T) (T, error) {
	curr := dll.Head
	for i := 0; i < dll.Length; i++ {
		if curr.Value == item {
			break
		}
		curr = curr.Next
	}

	if curr == nil {
		return *new(T), ErrItemNotFound
	}

	dll.Length--
	if dll.Length == 0 {
		dll.Head = nil
		dll.Tail = nil

		return curr.Value, nil
	}

	if curr.Prev != nil {
		curr.Prev = curr.Next
	}

	if curr.Next != nil {
		curr.Next = curr.Prev
	}

	if curr == dll.Head {
		dll.Head = curr.Next
	}

	if curr == dll.Tail {
		dll.Tail = curr.Prev
	}

	curr.Prev = nil
	curr.Next = nil

	return curr.Value, nil
}

func (dll *DoubleLinkedList[T]) Get(index int) (T, error) {
	return dll.GetAt(index)
}

func (dll *DoubleLinkedList[T]) GetAt(index int) (T, error) {
	_ = index
	// TODO: Implement.
	return *new(T), nil
}

func (dll *DoubleLinkedList[T]) RemoveAt(index int) (T, error) {
	// TODO: Implement.
	_ = index
	return *new(T), nil
}
