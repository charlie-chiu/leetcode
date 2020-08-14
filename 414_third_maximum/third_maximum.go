package third_maximum

import (
	"math"
	"sort"
)

var thirdMax func([]int) int = secondTry

func secondTry(nums []int) int {
	const min = math.MinInt64
	var first, second, third = min, min, min
	//maintain biggest 3 numbers
	for i := 0; i < len(nums); i++ {
		elem := nums[i]
		if elem > first {
			third = second
			second = first
			first = elem
		}

		if elem > second && elem < first {
			third = second
			second = elem
		}

		if elem > third && elem < second {
			third = elem
		}
	}

	// The third maximum does not exist, so the maximum is returned instead.
	if third == min {
		return first
	}

	return third
}

func firstTry(nums []int) int {

	// remove duplicate
	table := make(map[int]bool)
	var unique = []int{}
	for i := 0; i < len(nums); i++ {
		elem := nums[i]
		if _, ok := table[elem]; ok {
			continue
		}

		unique = append(unique, elem)
		table[elem] = true
	}

	// sort
	sort.Sort(sort.Reverse(sort.IntSlice(unique)))

	// The third maximum does not exist, so the maximum is returned instead.
	if len(unique) < 3 {
		return unique[0]
	}

	return unique[2]
}
