package container_with_most_water

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input  []int
	output int
}

func TestMaxArea(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{2, 1, 3}, 4},
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
	}

	for _, testCase := range TestCases {
		testName := fmt.Sprintf("test input %v got max area %d", testCase.input, testCase.output)
		t.Run(testName, func(t *testing.T) {
			got := maxArea(testCase.input)
			if got != testCase.output {
				t.Errorf("want %d, got %d", testCase.output, got)
			}
		})
	}
}
