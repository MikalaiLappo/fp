package fp

import (
	"testing"
	"strings"
)

func TestMap(t *testing.T) {
	input := strings.Split("aoeu aoeu aoeu", " ")
	transformer := func(s string) []string {
		return strings.Split(s, "o")
	}
	result := Map(input, transformer)

	if len(result) != 3 {
		t.Errorf("wrong output length of: %d", len(result))
	}

	// i'm too lazy and it's 5am
}
