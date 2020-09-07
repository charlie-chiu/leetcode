package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	//leetcode approach 3
	var lastAppeared = make(map[rune]int)
	var i, longest, currLength int

	input := []rune(s)
	for j := 0; j < len(input); j++ {
		currChar := input[j]
		appearedIndex, existed := lastAppeared[currChar]
		if existed && appearedIndex > i {
			i = appearedIndex
		}

		currLength = j - i + 1
		if currLength > longest {
			longest = currLength
		}

		lastAppeared[currChar] = j + 1
	}

	return longest
}
