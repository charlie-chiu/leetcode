package remove_duplicates_from_sorted_array

import (
	"testing"
)

type TestCase struct {
	input    []int
	length   int
	modified []int
}

func TestRemoveDuplicates(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1, 2}, 2, []int{1, 2}},
		{[]int{1, 1, 1, 2, 2, 3}, 5, []int{1, 1, 2, 2, 3}},
		{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}, 7, []int{0, 0, 1, 1, 2, 3, 3}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			var input = make([]int, len(tc.input))
			copy(input, tc.input)

			gotLength := removeDuplicates(tc.input)

			if gotLength != tc.length {
				t.Logf("wrong length on input %v", input)
				t.Logf("want: %v", tc.length)
				t.Logf(" got: %v", gotLength)
				t.Fail()
			}

			for i := 0; i < tc.length; i++ {
				if tc.input[i] != tc.modified[i] {
					t.Logf("wrong output since %dth element", i)
					t.Logf(" got: %v", tc.input)
					t.Logf("want: %v ...", tc.modified)
					t.Fail()
				}
			}
		})
	}
}
