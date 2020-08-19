package repeated_substring_pattern

import "strings"

var repeatedSubstringPattern = leetcode

func leetcode(s string) bool {
	// idea from leetcode discuss
	s2 := s + s
	s2 = s2[1 : len(s2)-1]

	return strings.Contains(s2, s)
}

// 8ms / 5MB
func firstTry(s string) bool {
	// use two pointers to indicate start and end of a substring
	// loop to check pattern
	// always start from 0
	const subFirst = 0
	var subLast = 0

	// for substring s[0:subLast]
	for subLast <= len(s)/2 {
		substring := s[subFirst : subLast+1]
		//log.Printf("substring : %v", substring)
		subLength := subLast - subFirst + 1

		patternStart := subLast + 1
		patternLast := patternStart + subLength - 1
		for patternLast < len(s) {
			//log.Println(patternStart, patternLast, pattern)
			if s[patternStart:patternLast+1] != substring {
				break
			}

			if patternLast == len(s)-1 {
				return true
			}
			patternStart += subLength
			patternLast += subLength
		}
		subLast++
	}

	return false
}
