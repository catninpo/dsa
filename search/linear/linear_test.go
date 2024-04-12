package linear_test

import (
	"testing"

	"github.com/catninpo/dsa/linear"
)

func TestLinearSearch(t *testing.T) {
	tt := map[string]struct {
		Haystack       []int
		Needle         int
		ExpectedResult bool
	}{
		"should find needle in haystack": {
			Haystack:       []int{1, 4, 3, 2, 7, 8},
			Needle:         7,
			ExpectedResult: true,
		},
		"should not find needle in haystack": {
			Haystack:       []int{1, 4, 3, 2, 7, 8},
			Needle:         155,
			ExpectedResult: false,
		},
	}

	for test, tc := range tt {
		t.Run(test, func(t *testing.T) {
			result := linear.Search(tc.Haystack, tc.Needle)
			if result != tc.ExpectedResult {
				t.Fatalf("[\u2717] expected result is not correct: got=%v want=%v",
					result, tc.ExpectedResult)
			}
			t.Logf("[\u2713] expected result is correct")
		})
	}
}
