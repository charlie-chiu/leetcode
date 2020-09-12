package trapping_rain_water

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input  []int
	answer int
}

func TestTrap(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{}, 0},
		{[]int{5}, 0},
		{[]int{1, 100}, 0},
		{[]int{1, 0, 1}, 1},
		{[]int{1, 0, 6}, 1},
		{[]int{6, 0, 2}, 2},
		{[]int{0, 2, 1, 0, 1, 3, 0}, 4},
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
	}

	for _, tc := range TestCases {
		t.Run(fmt.Sprintf("%v", tc.input), func(t *testing.T) {
			got := trap(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer on %v", tc.input)
				t.Errorf("want %v got %v", tc.answer, got)
			}
		})
	}
}
