package sort_colors

const (
	red   = 0
	white = 1
	blue  = 2
)

func sortColors(nums []int) {
	//DNF algorithm
	var L, M, R = 0, 0, len(nums) - 1

	for M <= R {
		switch nums[M] {
		case red:
			nums[L], nums[M] = nums[M], nums[L]
			L++
			M++
		case white:
			M++
		case blue:
			nums[M], nums[R] = nums[R], nums[M]
			R--
		}
	}
}
