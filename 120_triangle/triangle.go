package triangle

import (
	"sort"
)

// first try
// 0ms & 3.3 MB (39.64%)
func minimumTotal(triangle [][]int) int {
	var upperRowPaths = make([]int, 1)
	var currPath []int

	for _, nums := range triangle {
		currPath = make([]int, len(nums))
		for i, num := range nums {
			if i == 0 {
				// first
				currPath[i] = num + upperRowPaths[i]
			} else if i == len(nums)-1 {
				// last
				currPath[i] = num + upperRowPaths[i-1]
			} else {
				currPath[i] = num + min(upperRowPaths[i], upperRowPaths[i-1])
			}
		}
		upperRowPaths = currPath
	}

	sort.Ints(upperRowPaths)

	return upperRowPaths[0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
