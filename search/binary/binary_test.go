package binary_test

import (
	"testing"

	"github.com/catninpo/dsa/binary"
)

func TestBinarySearch(t *testing.T) {
	tt := map[string]struct {
		Haystack       []int
		Needle         int
		ExpectedResult int
	}{
		"should find needle in sorted haystack": {
			Haystack:       []int{1, 4, 8, 10, 15, 24},
			Needle:         15,
			ExpectedResult: 4,
		},
		"should not find needle in sorted haystack": {
			Haystack:       []int{1, 4, 8, 10, 15, 24},
			Needle:         6,
			ExpectedResult: -1,
		},
		"long haystack": {
			Haystack: []int{
				1,
				2,
				3,
				4,
				5,
				7,
				8,
				9,
				10,
				11,
				12,
				14,
				15,
				16,
				17,
				18,
				21,
				23,
				24,
			},
			Needle:         3,
			ExpectedResult: 2,
		},
	}

	for test, tc := range tt {
		t.Run(test, func(t *testing.T) {
			result := binary.Search(tc.Haystack, tc.Needle)
			if result != tc.ExpectedResult {
				t.Fatalf("[\u2717] imperative expected result is not correct: got=%v want=%v",
					result, tc.ExpectedResult)
			}
			t.Logf("[\u2713] imperative expected result is correct")

			result = binary.SearchRecursive(tc.Haystack, tc.Needle)
			if result != tc.ExpectedResult {
				t.Fatalf("[\u2717] recursive search expected result is not correct: got=%v want=%v",
					result, tc.ExpectedResult)
			}
			t.Logf("[\u2713] recursive expected result is correct")
		})
	}
}
