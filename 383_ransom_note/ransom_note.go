package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	var clipboard = make(map[rune]int)

	// cut letters from magazine
	for _, s := range magazine {

		clipboard[s]++
	}

	// paste to note
	for _, s := range ransomNote {
		if count, ok := clipboard[s]; !ok || count < 1 {
			return false
		}

		clipboard[s]--
	}

	return true
}
