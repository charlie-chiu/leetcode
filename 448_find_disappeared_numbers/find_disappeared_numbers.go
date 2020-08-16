package find_disappeared_numbers

import (
	"log"
)

var findDisappearedNumbers func([]int) []int = copycat

// copy from internet
// mark appeared number by change value to negative
func copycat(nums []int) []int {
	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	for _, num := range nums {
		markedIndex := abs(num) - 1

		nums[markedIndex] = abs(nums[markedIndex]) * -1
	}

	answer := []int{}
	for i, num := range nums {
		if num > 0 {
			answer = append(answer, i+1)
		}
	}

	return answer
}

func secondTry(nums []int) []int {
	answer := make([]int, len(nums)+1)
	for i := range answer {
		answer[i] = i
	}

	log.Printf("%+v\n", nums)
	//for _, num := range nums {
	//	answer[num] = -1
	//}

	//discard 0
	return answer[1:]
}

//using additional memory
func firstTry(nums []int) []int {
	answer := make([]int, 0)
	counter := make(map[int]int)
	for _, num := range nums {
		_, ok := counter[num]
		if ok {
			counter[num]++
		} else {
			counter[num] = 1
		}
	}

	//log.Println(counter)
	for i := 1; i <= len(nums); i++ {
		if _, ok := counter[i]; !ok {
			answer = append(answer, i)
		}
	}

	return answer
}
