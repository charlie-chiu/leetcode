package count_primes

import (
	"testing"
)

type TestCase struct {
	input, output int
}

func TestSolution(t *testing.T) {
	var TestCases = []TestCase{
		{0, 0},
		{4, 2}, // 2, 3
		{5, 2},
		{7, 3},
		{10, 4},
	}

	for _, testCase := range TestCases {
		got := solution(testCase.input)

		if got != testCase.output {
			t.Errorf("assert failed on input %d, want %d, got %d", testCase.input, testCase.output, got)
		}
	}
}

type IsPrimeTestCase struct {
	input  int
	output bool
}

func TestIsPrime(t *testing.T) {
	var TestCases = []IsPrimeTestCase{
		{1, false},
		{2, true},
		{3, true},
		{4, false},
		{5, true},
		{7, true},
		{10, false},
		{100, false},
		{349, true},
		{350, false},
	}

	for _, testCase := range TestCases {
		got := isPrime(testCase.input)
		if got != testCase.output {
			t.Errorf("assert failed on input %d, want %v, got %v", testCase.input, testCase.output, got)
		}
	}
}
