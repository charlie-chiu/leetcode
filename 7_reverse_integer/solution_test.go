package reverse_integer

import (
	"math"
	"testing"
)

type TestCase struct {
	input, output int
}

func TestSolution(t *testing.T) {

	testCases := []TestCase{
		{math.MaxInt32 + 1, 0},
		{math.MinInt32 - 1, 0},
		{123, 321},
		{123557, 755321},
		{120, 21},
		{-123, -321},
		{1534236469, 0},
	}

	for _, tc := range testCases {

		got := solution(tc.input)

		if got != tc.output {
			t.Errorf("assert %d reverse to %d, got %d", tc.input, tc.output, got)
		}
	}
}
