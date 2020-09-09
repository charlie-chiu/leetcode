package dutch_national_flag

// sort to 0,0,0,... 1,1,1...

const (
	zero = 0
	one  = 1
)

// using 2 pointers
// swap zero and one
func sort2Colors(nums []int) {
	var L, R = 0, len(nums) - 1

	for L < R {
		for nums[L] == zero && L < R {
			L++
		}
		for nums[R] == one && L < R {
			R--
		}

		if L < R {
			nums[L], nums[R] = nums[R], nums[L]
			L++
			R--
		}
	}
}

// by counting colors
//func sort2Colors(nums []int) {
//
//	var result = make([]int, 0)
//	var countZero int
//
//	for _, num := range nums {
//		if num == zero {
//			countZero++
//		}
//	}
//
//	for i := 0; i < countZero; i++ {
//		result = append(result, zero)
//	}
//	for i := countZero; i < len(nums); i++ {
//		result = append(result, one)
//	}
//
//	copy(nums, result)
//}
