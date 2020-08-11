package contains_duplicate

import "sort"

var solution func([]int) bool = sorting

//first try
// time complexity O(n)
func bruteForce(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func hashTable1(nums []int) bool {
	table := make(map[int]int, len(nums))
	for i, num := range nums {
		table[num] = i
	}

	for j, num := range nums {
		if i, ok := table[num]; ok && j != i {
			return true
		}
	}

	return false
}

func hashTable2(nums []int) bool {
	table := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := table[num]; ok && j != i {
			return true
		}

		table[num] = i
	}

	return false
}

func sorting(nums []int) bool {
	// leetcode solution: sorting
	sort.Ints(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}

	return false
}
