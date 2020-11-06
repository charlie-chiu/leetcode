package valid_perfect_square

import "testing"

type TestCase struct {
	num    int
	answer bool
}

func TestValidSquare(t *testing.T) {
	var TestCases = []TestCase{
		{16, true},
		{14, false},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := isPerfectSquare(tc.num)
			if got != tc.answer {
				t.Errorf("wrong answer on input %d, want %v got %v", tc.num, tc.answer, got)
			}
		})
	}
}
