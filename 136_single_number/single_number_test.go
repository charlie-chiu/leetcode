package single_number

import "testing"

type TestCase struct {
	input  []int
	output int
}

func TestSingleNumber(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{2, 2, 1}, 1},
		{[]int{3}, 3},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := singleNumber(tc.input)
			if got != tc.output {
				t.Log("wrong answer.\n")
				t.Logf("input: %v", tc.input)
				t.Logf("want:  %d", tc.output)
				t.Logf("got:   %d", got)
				t.Error()
			}
		})
	}
}
