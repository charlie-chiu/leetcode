package majority_element

import "testing"

type TestCase struct {
	input  []int
	answer int
}

func TestMajorityElement(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{0}, 0},
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := majorityElement(tc.input)
			if got != tc.answer {
				t.Logf("wrong answer at input %v\n", tc.input)
				t.Logf("want %d got %d", tc.answer, got)
				t.Fail()
			}
		})
	}
}
