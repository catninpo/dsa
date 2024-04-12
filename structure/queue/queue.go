package queue

type qNode[T any] struct {
	Value T
	Next  *qNode[T]
}

type Queue[T any] struct {
	Head *qNode[T]
	Tail *qNode[T]

	Length int
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
}

func (q *Queue[T]) Peek() T {
	if q.Head != nil {
		return q.Head.Value
	}

	return *new(T)
}

func (q *Queue[T]) Enqueue(v T) {
	q.Length++

	node := qNode[T]{
		Value: v,
		Next:  nil,
	}

	if q.Tail == nil {
		q.Head = &node
		q.Tail = &node
		return
	}

	q.Tail.Next = &node
	q.Tail = &node
}

func (q *Queue[T]) Dequeue() T {
	if q.Head == nil {
		return *new(T)
	}

	q.Length--

	head := q.Head
	q.Head = q.Head.Next

	if q.Length == 0 {
		q.Tail = nil
	}

	return head.Value
}
