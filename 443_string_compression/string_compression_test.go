package string_compression

import (
	"fmt"
	"reflect"
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
	}

	for _, testCase := range TestCases {
		t.Run(fmt.Sprintf("compress %s", testCase.input), func(t *testing.T) {
			length := compress(testCase.input)

			if !reflect.DeepEqual(testCase.input, testCase.output) {
				t.Errorf("wrong output.\nwant:%s\n got:%s", testCase.output, testCase.input)
			}

			if length != testCase.length {
				t.Errorf("wrong length.\nwant %d, got %d", testCase.length, length)
			}
		})
	}
}
