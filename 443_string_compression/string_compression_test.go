package string_compression

import (
	"fmt"
	"testing"
)

type TestCase struct {
	input, output []byte
	length        int
}

func TestName(t *testing.T) {
	var TestCases = []TestCase{
		{[]byte("a"), []byte("a"), 1},
		{[]byte("aabbccc"), []byte("a2b2c3"), 6},
		{[]byte("abbbbbbbbbbbb"), []byte("ab12"), 4},
	}

	for _, testCase := range TestCases {
		t.Run(fmt.Sprintf("compress %s", testCase.input), func(t *testing.T) {
			length := compress(testCase.input)

			if length != testCase.length {
				t.Logf("input is %s", testCase.input)
				t.Fatalf("wrong length.\nwant %d, got %d", testCase.length, length)
			}

			if len(testCase.input) < testCase.length {
				t.Logf("input is %s", testCase.input)
				t.Fatalf("input too short.\nexpected %d, got %d", testCase.length, len(testCase.input))
			}

			for i, b := range testCase.output {
				if testCase.input[i] != b {
					t.Errorf("wrong output.\nwant:%s\n got:%s", testCase.output, testCase.input)
				}
			}
		})
	}
}
