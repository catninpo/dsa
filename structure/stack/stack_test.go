package stack_test

import (
	"testing"

	"github.com/catninpo/dsa/structure/stack"
)

func TestNew(t *testing.T) {
	s := stack.New[int]()

	if s.Size != 0 {
		t.Fatalf("[\u2717] stack did not initialise to zero size: initialised_size=%d",
			s.Size)
	}
	t.Logf("[\u2713] stack initialised to correct size of zero")

	if s.Head != nil {
		t.Fatalf("[\u2717] stack has head set rather than nil: head=%v", s.Head)
	}
	t.Logf("[\u2713] stack has empty head as expected")
}

func TestPush(t *testing.T) {
	tt := map[string]struct {
		Values            []int
		ExpectedPeekValue int
		ExpectedSize      int
		ExpectedError     error
	}{
		"stack of size 3": {
			Values:            []int{5, 6, 3},
			ExpectedPeekValue: 3,
			ExpectedSize:      3,
			ExpectedError:     nil,
		},
		"stack of size 1": {
			Values:            []int{5},
			ExpectedPeekValue: 5,
			ExpectedSize:      1,
			ExpectedError:     nil,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := stack.New[int]()
			for _, v := range tc.Values {
				s.Push(v)
			}

			value, err := s.Peek()
			if err != tc.ExpectedError {
				t.Fatalf("[\u2717] unexpected error on peek: err=%v", err)
			}

			if value != tc.ExpectedPeekValue {
				t.Fatalf("[\u2717] unexpected value on peek: got=%d want=%d",
					value, tc.ExpectedPeekValue)
			}
			t.Logf("[\u2713] peek value of stack was correct")

			if s.Size != tc.ExpectedSize {
				t.Fatalf("[\u2717] unexpected size for stack: got=%d want=%d",
					s.Size, tc.ExpectedSize)
			}
			t.Logf("[\u2713] size of stack was correct")
		})
	}
}

func TestPop(t *testing.T) {
	tt := map[string]struct {
		Values           []int
		ExpectedPopValue int
		ExpectedSize     int
		ExpectedError    error
	}{
		"stack of size 3": {
			Values:           []int{5, 6, 3},
			ExpectedPopValue: 3,
			ExpectedSize:     2,
			ExpectedError:    nil,
		},
		"stack of size 1": {
			Values:           []int{5},
			ExpectedPopValue: 5,
			ExpectedSize:     0,
			ExpectedError:    nil,
		},
		"stack of size 0": {
			Values:           []int{},
			ExpectedPopValue: 0,
			ExpectedSize:     0,
			ExpectedError:    stack.ErrStackEmpty,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			s := stack.New[int]()
			for _, v := range tc.Values {
				s.Push(v)
			}

			value, err := s.Pop()
			if err != tc.ExpectedError {
				t.Fatalf("[\u2717] unexpected error on pop: err=%v", err)
			}

			if value != tc.ExpectedPopValue {
				t.Fatalf("[\u2717] unexpected value on pop: got=%d want=%d",
					value, tc.ExpectedPopValue)
			}
			t.Logf("[\u2713] pop value of stack was correct")

			if s.Size != tc.ExpectedSize {
				t.Fatalf("[\u2717] unexpected size for stack: got=%d want=%d",
					s.Size, tc.ExpectedSize)
			}
			t.Logf("[\u2713] size of stack was correct")
		})
	}
}
