package pascals_triangle

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input  int
	answer [][]int
}

func TestGenerate(t *testing.T) {
	var TestCases = []TestCase{
		{1, [][]int{
			{1},
		}},
		{3, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
		}},
		{5, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := generate(tc.input)
			if !reflect.DeepEqual(tc.answer, got) {
				t.Errorf("wrong answer")
				t.Logf("want %v", tc.answer)
				t.Logf(" got %v", got)
			}
		})
	}
}
