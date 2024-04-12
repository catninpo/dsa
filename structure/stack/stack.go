package stack

import "errors"

var ErrStackEmpty = errors.New("stack is empty")

type sNode[T any] struct {
	Value T
	Prev  *sNode[T]
}

type Stack[T any] struct {
	Head *sNode[T]

	Size int
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		Head: nil,
		Size: 0,
	}
}

func (s *Stack[T]) Push(item T) {
	s.Size++

	head := sNode[T]{
		Value: item,
		Prev:  s.Head,
	}

	s.Head = &head
}

func (s *Stack[T]) Pop() (T, error) {
	if s.Head == nil {
		return *new(T), ErrStackEmpty
	}

	s.Size--

	item := s.Head
	s.Head = s.Head.Prev

	return item.Value, nil
}

func (s *Stack[T]) Peek() (T, error) {
	if s.Head == nil {
		return *new(T), ErrStackEmpty
	}

	return s.Head.Value, nil
}
