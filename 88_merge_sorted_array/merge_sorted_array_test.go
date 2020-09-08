package merge_sorted_array

import (
	"reflect"
	"testing"
)

type TestCase struct {
	nums1  []int
	m      int
	nums2  []int
	n      int
	answer []int
}

func TestMerge(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{5, 0}, 1, []int{3}, 1, []int{3, 5}},
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}

	for _, tc := range TestCases {
		var nums1 = make([]int, len(tc.nums1))
		copy(nums1, tc.nums1)

		merge(tc.nums1, tc.m, tc.nums2, tc.n)

		if !reflect.DeepEqual(tc.nums1, tc.answer) {
			t.Errorf("wrong output.")
			t.Logf("nums1: %v", nums1)
			t.Logf("nums2: %v", tc.nums2)
			t.Logf("  got: %v", tc.nums1)
			t.Logf(" want: %v", tc.answer)

		}
	}
}
