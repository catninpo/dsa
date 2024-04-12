package twocrystalballs_test

import (
	"testing"

	"github.com/catninpo/dsa/problems/twocrystalballs"
)

func TestSolution(t *testing.T) {
	tt := map[string]struct {
		Input          []bool
		ExpectedResult int
	}{
		"returns correct index when one is present": {
			Input:          []bool{false, false, false, true, true, true},
			ExpectedResult: 3,
		},
		"returns -1 when the balls never break": {
			Input:          []bool{false, false, false, false},
			ExpectedResult: -1,
		},
	}

	for test, tc := range tt {
		t.Run(test, func(t *testing.T) {
			result := twocrystalballs.Solve(tc.Input)
			if result != tc.ExpectedResult {
				t.Fatalf("[\u2717] expected result is not correct: got=%v want=%v",
					result, tc.ExpectedResult)
			}
			t.Logf("[\u2713] expected result is correct")
		})
	}
}
