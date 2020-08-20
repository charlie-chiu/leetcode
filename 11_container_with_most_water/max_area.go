package container_with_most_water

// 12ms 5.9MB
func maxArea(height []int) int {
	var min = func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var max = func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	var getArea = func(height []int, l int, r int) int {
		return min(height[l], height[r]) * (r - l)
	}

	var highestL int
	var highestR int
	var maxArea int
	var l = 0
	var r = len(height) - 1

	maxArea = getArea(height, l, r)
	for l < r {
		//move shorter bar
		if height[l] < height[r] {
			l++
		} else {
			r--
		}

		if height[l] > highestL {
			highestL = height[l]
			maxArea = max(maxArea, getArea(height, l, r))
		}
		if height[r] > highestR {
			highestR = height[r]
			maxArea = max(maxArea, getArea(height, l, r))
		}
	}

	return maxArea
}
