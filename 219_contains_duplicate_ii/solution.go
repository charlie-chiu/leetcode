package contains_duplicate

var solution func([]int, int) bool = hashTable2

func bruteForce(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if j-i > k {
				break
			}

			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

func hashTable(nums []int, k int) bool {
	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	// space complexity : use additional space
	table := make(map[int][]int)
	for index, num := range nums {
		table[num] = append(table[num], index)
	}

	for index := 0; index < len(nums); index++ {
		num := nums[index]
		if indices, ok := table[num]; ok {
			for _, secondIndex := range indices {
				if index != secondIndex && abs(secondIndex-index) <= k {
					return true
				}
			}
		}
	}

	return false
}

func hashTable2(nums []int, k int) bool {
	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	// space complexity : use additional space
	table := make(map[int]map[int]bool, len(nums))
	for index, num := range nums {
		if table[num] == nil {
			table[num] = make(map[int]bool)
		}

		for existIndex := range table[num] {
			if abs(index-existIndex) <= k {
				return true
			}
		}

		table[num][index] = true
	}

	return false
}
