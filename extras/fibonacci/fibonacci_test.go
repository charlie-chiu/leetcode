package fibonacci

import (
	"testing"
)

type TestCase struct {
	input, answer int
}

func TestFibonacci(t *testing.T) {
	var TestCases = []TestCase{
		{0, 0},
		{1, 1},
		{5, 5},
		{10, 55},
		{12, 144},
		{20, 6765},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := fibonacci(tc.input)
			if got != tc.answer {
				t.Errorf("expected %dth number is %d, got %d", tc.input, tc.answer, got)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(100)
	}
}
