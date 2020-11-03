package triangle

import (
	"sort"
)

//second try, modify original slice
func minimumTotal(triangle [][]int) int {
	for row := 1; row < len(triangle); row++ {
		nums := triangle[row]
		for i, num := range nums {
			if i == 0 {
				triangle[row][i] = num + triangle[row-1][i]
			} else if i == len(nums)-1 {
				triangle[row][i] = num + triangle[row-1][i-1]
			} else {
				triangle[row][i] = num + min(triangle[row-1][i], triangle[row-1][i-1])
			}
		}
	}

	sort.Ints(triangle[len(triangle)-1])

	return triangle[len(triangle)-1][0]
}

// first try
// 0ms & 3.3 MB (39.64%)
//func minimumTotal(triangle [][]int) int {
//	var upperRowPaths = make([]int, 1)
//	var currPath []int
//
//	for _, nums := range triangle {
//		currPath = make([]int, len(nums))
//		for i, num := range nums {
//			if i == 0 {
//				// first
//				currPath[i] = num + upperRowPaths[i]
//			} else if i == len(nums)-1 {
//				// last
//				currPath[i] = num + upperRowPaths[i-1]
//			} else {
//				currPath[i] = num + min(upperRowPaths[i], upperRowPaths[i-1])
//			}
//		}
//		upperRowPaths = currPath
//	}
//
//	sort.Ints(upperRowPaths)
//
//	return upperRowPaths[0]
//}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
