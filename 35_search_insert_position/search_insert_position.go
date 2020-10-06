package search_insert_position

// from leetcode discuss
// binary search using two points
func searchInsert(nums []int, target int) int {
	var low = 0
	var high = len(nums) - 1
	var mid int
	for low <= high {
		mid = low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}

//brute force, in linear time
// 4ms (81%) on leetcode
//func searchInsert(nums []int, target int) int {
//	for i, num := range nums {
//		if num >= target {
//			return i
//		}
//	}
//	return len(nums)
//}
