package first_unique_character

import (
	"strings"
)

var firstUniqChar = third

// leetcode sample
// 8ms / 5.7 MB
func third(s string) int {
	var counter [26]int
	for _, char := range s {
		counter[char-'a']++
	}

	for i, char := range s {
		if counter[char-'a'] == 1 {
			return i
		}
	}

	return -1
}

// 112ms / 5.8MB
func secondTry(s string) int {
	for i, c := range s {
		if strings.Count(s, string(c)) == 1 {
			return i
		}
	}
	return -1
}

// 20ms / 6.7 MB
func firstTry(s string) int {
	counter := make(map[rune]int, len(s))

	for _, char := range s {
		counter[char]++
	}

	for i, char := range s {
		if counter[char] == 1 {
			return i
		}
	}

	return -1
}
