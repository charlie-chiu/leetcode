package group_anagrams

import (
	"reflect"
)

// todo: TLE on leetcode
// todo: test needs to improve (answer order doesn't matter)

func groupAnagrams(strs []string) [][]string {
	var letterCount map[rune]int
	var countGroup []map[rune]int
	var resultGroup = make([][]string, 0)

	for _, str := range strs {
		letterCount = map[rune]int{}
		for _, letter := range str {
			letterCount[letter] += 1
		}

		index, ok := findIndex(countGroup, letterCount)
		if ok {
			resultGroup[index] = append(resultGroup[index], str)
		} else {
			countGroup = append(countGroup, letterCount)
			resultGroup = append(resultGroup, []string{str})
		}
	}

	return resultGroup
}

func findIndex(from []map[rune]int, target map[rune]int) (index int, ok bool) {
	for i := 0; i < len(from); i++ {
		if reflect.DeepEqual(target, from[i]) {
			return i, true
		}
	}
	return -1, false
}
