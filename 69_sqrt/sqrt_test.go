package sqrt

import "testing"

type TestCase struct {
	input, answer int
}

func TestSqrt(t *testing.T) {
	var TestCases = []TestCase{
		{4, 2},
		{8, 2},
		{9, 3},
		{630, 25},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := mySqrt(tc.input)
			if got != tc.answer {
				t.Errorf("got %d want %d", got, tc.answer)
			}
		})
	}
}
