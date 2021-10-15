package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counterS := make(map[byte]int, 0)
	for i := 0; i < len(s); i++ {
		counterS[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if count, ok := counterS[t[i]]; ok && count > 0 {
			counterS[t[i]]--
		} else {
			return false
		}
	}

	return true
}
