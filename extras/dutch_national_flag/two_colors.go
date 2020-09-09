package dutch_national_flag

// sort to 0,0,0,... 1,1,1...

const (
	zero = 0
	one  = 1
)

// by counting colors
func sort2Colors(nums []int) {

	var result = make([]int, 0)
	var countZero int

	for _, num := range nums {
		if num == zero {
			countZero++
		}
	}

	for i := 0; i < countZero; i++ {
		result = append(result, zero)
	}
	for i := countZero; i < len(nums); i++ {
		result = append(result, one)
	}

	copy(nums, result)
}
