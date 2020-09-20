package two_sum

var solution func([]int, int) []int = onePassHashTable

// first try time complexity: O(n^2)
func bruteForce(nums []int, target int) []int {
	for firstIndex, num := range nums {
		for i2, num2 := range nums[firstIndex+1:] {
			if num+num2 == target {
				return []int{firstIndex, i2 + firstIndex + 1}
			}
		}
	}

	return nil
}

//from leetcode solution,
//time complexity: O(n)
//space complexity: O(n)
func twoPassHashTable(nums []int, target int) []int {
	table := make(map[int]int, len(nums))
	for i, num := range nums {
		table[num] = i
	}

	for i, num := range nums {
		t := target - num

		if i2, ok := table[t]; ok && i != i2 {
			return []int{i, i2}
		}
	}

	return nil
}

//from leetcode solution,
//time complexity: O(n)
//space complexity: O(n)
func onePassHashTable(nums []int, target int) []int {
	table := make(map[int]int, len(nums))
	for i, num := range nums {
		complement := target - num

		if i2, ok := table[complement]; ok {
			return []int{i2, i}
		}

		table[num] = i
	}

	return nil
}
