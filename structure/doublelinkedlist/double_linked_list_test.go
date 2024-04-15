package doublelinkedlist_test

import (
	"testing"

	"github.com/catninpo/dsa/structure/doublelinkedlist"
)

func TestNew(t *testing.T) {
	dll := doublelinkedlist.New[int]()
	if dll.Length != 0 {
		t.Fatalf("[\u2717] length should be zero: got=%d", dll.Length)
	}
	t.Logf("[\u2713] length was set to zero on New()")

	if dll.Head != nil {
		t.Fatalf("[\u2717] Head should be nil: got=%v", dll.Head)
	}
	t.Logf("[\u2713] list head was set to nil on New()")
}

func TestPrepend(t *testing.T) {
	tt := map[string]struct {
		StartingList []int
		ExpectedList []int
		Elements     []int
	}{
		"small starting list": {
			StartingList: []int{1, 2},
			ExpectedList: []int{3, 1, 2},
			Elements:     []int{3},
		},
		"empty starting list": {
			StartingList: []int{},
			ExpectedList: []int{1},
			Elements:     []int{1},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			dll := doublelinkedlist.New[int]()
			for _, item := range tc.StartingList {
				dll.Append(item)
			}

			for _, item := range tc.Elements {
				dll.Prepend(item)
			}

			curr := dll.Head
			for i := 0; i < dll.Length; i++ {
				if curr.Value != tc.ExpectedList[i] {
					t.Fatalf("[\u2717] list values did not match: index=%d got=%d want=%d",
						i, curr.Value, tc.ExpectedList[i])
				}
				curr = curr.Next
			}
		})
	}
}

func TestAppend(t *testing.T) {
	tt := map[string]struct {
		ExpectedList []int
		Elements     []int
	}{
		"empty list": {
			ExpectedList: []int{3},
			Elements:     []int{3},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			dll := doublelinkedlist.New[int]()
			for _, item := range tc.ExpectedList {
				dll.Append(item)
			}

			curr := dll.Head
			for i := 0; i < dll.Length; i++ {
				if curr.Value != tc.ExpectedList[i] {
					t.Fatalf("[\u2717] list values did not match: index=%d got=%d want=%d",
						i, curr.Value, tc.ExpectedList[i])
				}
				curr = curr.Next
			}
		})
	}
}
