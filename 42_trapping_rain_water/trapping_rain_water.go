package trapping_rain_water

func trap(height []int) int {
	//brute force
	width := len(height)
	if width < 3 {
		return 0
	}

	var water int

	for i, h := range height {
		// find water for every bar in heights

		highestL := 0
		for L := i - 1; L >= 0; L-- {
			highestL = max(highestL, height[L])
		}

		highestR := 0
		for R := i + 1; R < width; R++ {
			highestR = max(highestR, height[R])
		}

		water += max(min(highestL, highestR)-h, 0)

	}

	return water
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
