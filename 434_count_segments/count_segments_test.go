package count_segments

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input    string
	solution int
}

func TestCountSegments(t *testing.T) {
	var TestCases = []TestCase{
		{"", 0},
		{" ", 0},
		{"gopher", 1},
		{"Hello World", 2},
		{"Hello            World", 2},
		{"Hello, my name is John", 5},
	}

	for _, testCase := range TestCases {
		t.Run(fmt.Sprintf("count %s", testCase.input), func(t *testing.T) {
			got := countSegments(testCase.input)
			if got != testCase.solution {
				t.Errorf("wrong answer,want %d,  got %d", testCase.solution, got)
			}
		})
	}
}
