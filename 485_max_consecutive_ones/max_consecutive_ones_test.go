package max_consecutive_ones

import "testing"

type TestCase struct {
	input  []int
	output int
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	TestCases := []TestCase{
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, 0}, 2},
		{[]int{0, 1, 1}, 2},
		{[]int{1, 1, 0, 1, 1, 1}, 3},
		{[]int{1, 0, 1, 1, 0, 1}, 2},
	}

	for _, tc := range TestCases {
		got := findMaxConsecutiveOnes(tc.input)

		if got != tc.output {
			t.Errorf("wrong answer:\ninput: %v\n want: %v\n  got: %v", tc.input, tc.output, got)
		}
	}
}
