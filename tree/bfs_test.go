package tree_test

import (
	"testing"

	"github.com/catninpo/dsa/tree"
)

func TestBFS(t *testing.T) {
	tt := map[string]struct {
		Needle         int
		ExpectedResult bool
	}{
		"top level existing needle can be found": {
			Needle:         5,
			ExpectedResult: true,
		},
		"needle does not exist in list": {
			Needle:         904,
			ExpectedResult: false,
		},
		"needle exists in bottom mode": {
			Needle:         16,
			ExpectedResult: true,
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			found := tree.BFS(newBinaryTree(), tc.Needle)
			if found != tc.ExpectedResult {
				t.Fatalf("[\u2717] expected found result not correct: got=%v want=%v",
					found, tc.ExpectedResult)
			}
			t.Logf("[\u2713] expected found was was correct")
		})
	}
}
