package house_robber

import "testing"

type TestCase struct {
	input  []int
	answer int
}

func TestRob(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{114, 117, 207, 117, 235, 82, 90, 67, 143, 146, 53, 108, 200, 91, 80, 223, 58, 170, 110, 236, 81, 90, 222, 160, 165, 195, 187, 199, 114, 235, 197, 187, 69, 129, 64, 214, 228, 78, 188, 67, 205, 94, 205, 169, 241, 202, 144, 240}, 4173},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := rob(tc.input)
			if got != tc.answer {
				t.Errorf("wrong answer on input %v, want %d got %d", tc.input, tc.answer, got)
			}
		})
	}
}
