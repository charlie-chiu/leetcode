package insert_interval

import "sort"

// another approach
//func insert(intervals [][]int, newInterval []int) [][]int {
//	var i int
//	result := make([][]int, 0)
//	for ;i < len(intervals); i++ {
//		curr := intervals[i]
//
//		if curr[1] < newInterval[0] {
//			result = append(result, curr)
//		} else if newInterval[1] < curr[0] {
//			break
//		} else {
//			newInterval[0] = min(newInterval[0], curr[0])
//			newInterval[1] = max(newInterval[1], curr[1])
//		}
//	}
//
//	return append(append(result, newInterval), intervals[i:]...)
//}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)

	if len(intervals) == 1 {
		return intervals
	}

	// same as 56 merge intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	result = append(result, intervals[0])
	for curr := 1; curr < len(intervals); curr++ {
		// start from second interval
		currInterval := intervals[curr]
		lastInterval := result[len(result)-1]

		if !isOverlapping(currInterval, lastInterval) {
			result = append(result, currInterval)
		} else {
			result[len(result)-1][0] = min(currInterval[0], lastInterval[0])
			result[len(result)-1][1] = max(currInterval[1], lastInterval[1])
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func isOverlapping(a, b []int) bool {
	if a[0] < b[0] && a[1] < b[0] {
		return false
	}

	if b[0] < a[0] && b[1] < a[0] {
		return false
	}

	return true
}
