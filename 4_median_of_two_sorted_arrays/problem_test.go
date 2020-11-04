package median_of_two_sorted_arrays

import "testing"

type TestCase struct {
	nums1, nums2 []int
	answer       float64
}

func TestApproach(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0.0},
		{[]int{}, []int{1}, 1.0},
		{[]int{}, []int{2}, 2},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := findMedianSortedArrays(tc.nums1, tc.nums2)
			if got != tc.answer {
				t.Logf("wrong answer")
				t.Logf("nums1 : %v", tc.nums1)
				t.Logf("nums2 : %v", tc.nums2)
				t.Logf("want  : %v", tc.answer)
				t.Logf(" got  : %v", got)
				t.Fail()
			}
		})
	}
}
