package remove_element

import "testing"

type TestCase struct {
	nums   []int
	val    int
	length int
	//note : the order of those five elements can be arbitrary.
	modified []int
}

func TestRemoveDuplicates(t *testing.T) {
	var TestCases = []TestCase{
		{[]int{}, 3, 0, []int{}},
		{[]int{3, 2, 2, 3}, 3, 2, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5, []int{0, 1, 3, 0, 4}},
	}

	for _, tc := range TestCases {
		var input = make([]int, len(tc.nums))
		copy(input, tc.nums)

		gotLength := removeElement(tc.nums, tc.val)

		if gotLength != tc.length {
			t.Errorf("wrong length on input %v", input)
			t.Logf("want %d got %d", tc.length, gotLength)
			t.FailNow()
		}

		for i := 0; i < tc.length; i++ {
			if tc.nums[i] != tc.modified[i] {
				t.Logf("wrong output since %dth element", i)
				t.Logf(" got: %v", tc.nums)
				t.Logf("want: %v ...", tc.modified)
				t.FailNow()
			}
		}

	}
}
