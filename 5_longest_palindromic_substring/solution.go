package longest_palindromic_substring

//12ms beats 58.96%
func longestPalindrome(s string) string {
	var longestSubstring string

	for i := range s {
		var start, end = i, i
		var subStringLength int = 0

		for start >= 0 && end <= len(s)-1 {
			subString := s[start : end+1]
			if s[start] == s[end] {
				subStringLength = end - start + 1
				if subStringLength > len(longestSubstring) {
					longestSubstring = subString
				}
				start--
				end++
			} else {
				break
			}
		}
	}

	for i := range s {
		var start, end = i, i
		var subStringLength int = 0

		// even number string
		if end < len(s)-1 && s[start] == s[end+1] {
			end++
		} else {
			continue
		}

		for start >= 0 && end <= len(s)-1 {
			subString := s[start : end+1]
			if s[start] == s[end] {
				subStringLength = end - start + 1
				if subStringLength > len(longestSubstring) {
					longestSubstring = subString
				}
				start--
				end++
			} else {
				break
			}
		}
	}

	return longestSubstring
}

//brute force
// 872 ms
//func longestPalindrome(s string) string {
//	var longestSubstring string
//
//	for start := 0; start < len(s); start++ {
//		for end := start; end < len(s); end++ {
//			//log.Println(start, end)
//			subString := s[start : end+1]
//			//log.Println(subString)
//			if isPalindrome(subString) && len(subString) > len(longestSubstring) {
//				longestSubstring = subString
//			}
//		}
//	}
//
//	return longestSubstring
//}
//
//func isPalindrome(s string) bool {
//	length := len(s)
//	head := 0
//	tail := length - 1
//
//	for head < length/2 {
//		if s[head] != s[tail] {
//			return false
//		}
//		head++
//		tail--
//	}
//
//	return true
//}
