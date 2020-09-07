package longest_substring_without_repeating_characters

// leetcode solution
func lengthOfLongestSubstring(s string) int {
	var existed = make(map[rune]bool)
	var longest, currLength int

	for i, _ := range s {
		existed = map[rune]bool{}
		for j := i; j < len(s); j++ {
			//character repeated
			currChar := []rune(s)[j]
			currLength = j - i + 1
			if _, ok := existed[currChar]; ok {
				currLength--
				break
			}

			existed[currChar] = true
		}

		if currLength > longest {
			longest = currLength
		}
	}

	return longest
}
