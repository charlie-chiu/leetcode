package merge_intervals

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input, answer [][]int
}

func TestMerge(t *testing.T) {
	var TestCases = []TestCase{
		{
			[][]int{{1, 2}, {4, 5}},
			[][]int{{1, 2}, {4, 5}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}},
			[][]int{{1, 10}},
		},
		{
			[][]int{{1, 4}, {0, 4}},
			[][]int{{0, 4}},
		},
		{
			[][]int{{1, 4}, {0, 1}},
			[][]int{{0, 4}},
		},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := merge(tc.input)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Errorf("wrong answer on input %v", tc.input)
				t.Errorf("got %v", got)
				t.Fail()
			}
		})
	}
}

type overlappingTC struct {
	a, b     []int //interval
	expected bool
}

func TestIsOverlapping(t *testing.T) {
	testcases := []overlappingTC{
		// a -> b
		{[]int{0, 1}, []int{3, 4}, false},
		{[]int{0, 3}, []int{2, 4}, true},
		{[]int{0, 5}, []int{2, 4}, true},
		// b -> a
		{[]int{2, 7}, []int{0, 5}, true},
		{[]int{3, 4}, []int{0, 5}, true},
		{[]int{7, 9}, []int{0, 5}, false},

		// completed overlapping
		{[]int{3, 4}, []int{3, 4}, true},
	}

	for _, tc := range testcases {
		got := isOverlapping(tc.a, tc.b)
		if got != tc.expected {
			t.Logf("a: %v, b: %v", tc.a, tc.b)
			t.Logf("want %v, got %v", tc.expected, got)
			t.Fail()
		}
	}
}
