package queue_test

import (
	"testing"

	"github.com/catninpo/dsa/structure/queue"
)

func TestNew(t *testing.T) {
	q := queue.New[int]()

	if q.Head != nil {
		t.Fatalf("[\u2717] queue head was not initialised to nil")
	}
	t.Logf("[\u2713] queue head was initialised to nil")

	if q.Tail != nil {
		t.Fatalf("[\u2717] queue tail not was initialised to nil")
	}
	t.Logf("[\u2713] queue tail was initialised to nil")
}

func TestPeek(t *testing.T) {
	tt := map[string]struct {
		QueueValues        []int
		ExpectedPeekResult int
	}{
		"single item on queue returns value": {
			QueueValues:        []int{5},
			ExpectedPeekResult: 5,
		},
		"no items in queue returns zero value": {
			QueueValues:        []int{},
			ExpectedPeekResult: 0,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			q := queue.New[int]()
			for _, v := range tc.QueueValues {
				q.Enqueue(v)
			}

			val := q.Peek()
			if val != tc.ExpectedPeekResult {
				t.Fatalf("[\u2717] expected value was not correct: got=%d want=%d",
					val, tc.ExpectedPeekResult)
			}
			t.Logf("[\u2713] expected value was correct: got=%d want=%d",
				val, tc.ExpectedPeekResult)
		})
	}
}

func TestEnqueue(t *testing.T) {
	tt := map[string]struct {
		QueueValues        []int
		ExpectedPeekResult int
		ExpectedLength     int
	}{
		"push one item onto queue": {
			QueueValues:        []int{5},
			ExpectedPeekResult: 5,
			ExpectedLength:     1,
		},
		"push multiple items onto queue": {
			QueueValues:        []int{3, 7, 6},
			ExpectedPeekResult: 3,
			ExpectedLength:     3,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			q := queue.New[int]()
			for _, v := range tc.QueueValues {
				q.Enqueue(v)
			}

			if q.Peek() != tc.ExpectedPeekResult {
				t.Fatalf("[\u2717] expected peek value was not correct: got=%d want=%d",
					q.Peek(), tc.ExpectedPeekResult)
			}
			t.Logf("[\u2713] expected peek value was correct: got=%d want=%d",
				q.Peek(), tc.ExpectedPeekResult)

			if q.Length != tc.ExpectedLength {
				t.Fatalf("[\u2717] expected queue length was not correct: got=%d want=%d",
					q.Length, tc.ExpectedLength)
			}
			t.Logf("[\u2713] expected queue length was correct: got=%d want=%d",
				q.Length, tc.ExpectedLength)
		})
	}
}

func TestDequeue(t *testing.T) {
	tt := map[string]struct {
		QueueValues        []int
		ExpectedResult     int
		ExpectedPeekResult int
	}{
		"single item on queue returns zero value": {
			QueueValues:        []int{5},
			ExpectedResult:     5,
			ExpectedPeekResult: 0,
		},
		"no items in queue returns zero value": {
			QueueValues:        []int{},
			ExpectedResult:     0,
			ExpectedPeekResult: 0,
		},
		"two items in queue returns second value": {
			QueueValues:        []int{5, 3},
			ExpectedResult:     5,
			ExpectedPeekResult: 3,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			q := queue.New[int]()
			for _, v := range tc.QueueValues {
				q.Enqueue(v)
			}

			val := q.Dequeue()
			if val != tc.ExpectedResult {
				t.Fatalf("[\u2717] expected value was not correct: got=%d want=%d",
					val, tc.ExpectedResult)
			}
			t.Logf("[\u2713] expected value was correct: got=%d want=%d",
				val, tc.ExpectedResult)

			if tc.ExpectedPeekResult != q.Peek() {
				t.Fatalf("[\u2717] expected peek was not correct: got=%d want=%d",
					q.Peek(), tc.ExpectedPeekResult)
			}
			t.Logf("[\u2713] expected peek result was correct: got=%d want=%d",
				q.Peek(), tc.ExpectedPeekResult)
		})
	}
}
