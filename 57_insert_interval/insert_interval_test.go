package insert_interval

import (
	"reflect"
	"testing"
)

type TestCase struct {
	intervals   [][]int
	newInterval []int
	answer      [][]int
}

func TestInsert(t *testing.T) {
	var TestCases = []TestCase{
		{
			[][]int{},
			[]int{2, 5},
			[][]int{{2, 5}},
		},
		{
			[][]int{{1, 5}},
			[]int{2, 3},
			[][]int{{1, 5}},
		},
		{
			[][]int{{1, 5}},
			[]int{6, 8},
			[][]int{{1, 5}, {6, 8}},
		},
		{
			[][]int{{1, 5}},
			[]int{2, 7},
			[][]int{{1, 7}},
		},
		{
			[][]int{{1, 3}, {6, 9}},
			[]int{2, 5},
			[][]int{{1, 5}, {6, 9}},
		},
		{
			[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			[]int{4, 8},
			[][]int{{1, 2}, {3, 10}, {12, 16}},
		},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := insert(tc.intervals, tc.newInterval)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Errorf("wrong answer on input %v, %v", tc.intervals, tc.newInterval)
				t.Errorf("want %v", tc.answer)
				t.Errorf(" got %v", got)
				t.Fail()
			}
		})
	}
}
