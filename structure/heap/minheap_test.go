package heap_test

import (
	"testing"

	"github.com/catninpo/dsa/structure/heap"
)

func TestNew(t *testing.T) {
	h := heap.NewMinHeap[int]()

	if h.Length != 0 {
		t.Fatalf("[\u2717] min heap did not set length to 0 on initialise")
	}

	if len(h.Data) != 0 {
		t.Fatalf("[\u2717] min heap did not set empty data array on initialise")
	}
}
