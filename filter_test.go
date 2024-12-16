package fp

import "testing"

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	predicate := func(_ int, x int) bool { return x%2 == 0 }
	result := Filter(input, predicate)

	if len(result) != 2 {
		t.Errorf("wrong output length of: %d", len(result))
	}

	if result[0] != 2 && result[1] != 4 {
		t.Errorf("invalid output, should be [2, 4], got: %v", result)
	}
}
