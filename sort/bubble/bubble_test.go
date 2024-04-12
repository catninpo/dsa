package bubble_test

import (
	"reflect"
	"testing"

	"github.com/catninpo/dsa/sort/bubble"
)

func TestSort(t *testing.T) {
	tt := map[string]struct {
		Input  []int
		Output []int
	}{
		"small array sort": {
			Input:  []int{3, 1, 4, 2, 7},
			Output: []int{1, 2, 3, 4, 7},
		},
		"small array already sorted": {
			Input:  []int{1, 2, 4, 5, 6, 7, 8},
			Output: []int{1, 2, 4, 5, 6, 7, 8},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			bubble.Sort(tc.Input)

			if !reflect.DeepEqual(tc.Output, tc.Input) {
				t.Fatalf("[\u2717] post-sorted array did not match expected: got=%v want=%v",
					tc.Output, tc.Input)
			}
			t.Logf("[\u2713] post-sorted array matched expected value")
		})
	}
}
