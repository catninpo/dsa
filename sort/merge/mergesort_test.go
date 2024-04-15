package mergesort_test

import (
	"reflect"
	"testing"

	mergesort "github.com/catninpo/dsa/sort/merge"
)

func TestSort(t *testing.T) {
	tt := map[string]struct {
		Input  []int
		Output []int
	}{
		"small array with duplicate numbers": {
			Input:  []int{3, 2, 6, 4, 1, 4, 9},
			Output: []int{1, 2, 3, 4, 4, 6, 9},
		},
		"small array": {
			Input:  []int{4, 2, 3, 1},
			Output: []int{1, 2, 3, 4},
		},
		"middle pivot": {
			Input:  []int{4, 2, 5, 1, 3},
			Output: []int{1, 2, 3, 4, 5},
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			output := mergesort.Sort(tc.Input)

			if !reflect.DeepEqual(tc.Output, output) {
				t.Fatalf("[\u2717] post-sorted array did not match expected: got=%v want=%v",
					tc.Input, output)
			}
			t.Logf("[\u2713] post-sorted array matched expected value")
		})
	}
}
