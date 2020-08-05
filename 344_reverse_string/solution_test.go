package reverse_string

import (
	"bytes"
	"testing"
)

type TestCase struct {
	input, expect string
}

func TestSolution(t *testing.T) {
	testCases := []TestCase{
		{"hello", "olleh"},
		{"marmot", "tomram"},
	}

	for _, tc := range testCases {
		input := []byte(tc.input)
		expect := []byte(tc.expect)
		solution(input)

		if bytes.Compare(input, expect) != 0 {
			t.Errorf("not equal, want %q got %q", expect, input)
		}
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution([]byte("Hello"))
	}
}
