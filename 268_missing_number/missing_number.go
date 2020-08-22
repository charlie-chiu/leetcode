package missing_number

import (
	"sort"
)

var solution = try3

func try3(nums []int) int {
	//find by xor
	var result = len(nums)

	for i := 0; i < len(nums); i++ {
		result ^= i
		result ^= nums[i]
	}

	return result
}
func try2(nums []int) int {
	//find by sum
	sum := len(nums)

	for i, num := range nums {
		sum += i
		sum -= num
	}

	return sum
}

//24ms // 6MB
func try1(nums []int) int {
	sort.Ints(nums)

	length := len(nums)
	for i := 0; i <= length; i++ {
		if i == length || i != nums[i] {
			return i
		}
	}

	return 0
}
