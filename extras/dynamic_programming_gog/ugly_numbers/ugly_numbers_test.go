package ugly_numbers

import "testing"

type IsUglyTestCase struct {
	input  int
	answer bool
}

func TestIsUgly(t *testing.T) {
	var TestCases = []IsUglyTestCase{
		{0, false},
		{1, true},
		{3, true},
		{5, true},
		{7, false},
		{10, true},
		{11, false},
		{13, false},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := isUgly(tc.input)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("isUgly(%d) want %v got %v", tc.input, tc.answer, got)
				t.Fail()
			}
		})
	}
}

type GetNthTestCase struct {
	n, answer int
}

func TestGetUglyNumber(t *testing.T) {
	var TestCases = []GetNthTestCase{
		{1, 1},
		{2, 2},
		{3, 3},
		{7, 8},
		{10, 12},
		{15, 24},
		{150, 5832},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := getNthUglyNumber(tc.n)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("%dth ugly number is %v, got %v", tc.n, tc.answer, got)
				t.Fail()
			}
		})
	}
}
