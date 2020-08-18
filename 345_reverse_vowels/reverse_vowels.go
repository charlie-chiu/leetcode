package reverse_vowels

var reverseVowels = secondTry

func secondTry(s string) string {
	var isVowel = func(r rune) bool {
		switch r {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			return true
		}

		return false
	}

	var leftP = 0
	var rightP = len(s) - 1
	var inputRune = []rune(s)
	for leftP < rightP {
		if !isVowel(inputRune[leftP]) {
			leftP++
		} else if !isVowel(inputRune[rightP]) {
			rightP--
		} else {
			inputRune[leftP], inputRune[rightP] = inputRune[rightP], inputRune[leftP]
			leftP++
			rightP--
		}
	}

	return string(inputRune)
}

func firstTry(s string) string {
	var isVowel = func(r rune) bool {
		vowels := map[rune]bool{
			97:  true, //a
			101: true, //e
			105: true, //i
			111: true, //o
			117: true, //u
			65:  true, //A
			69:  true, //E
			73:  true, //I
			79:  true, //O
			85:  true, //U
		}
		_, ok := vowels[r]
		return ok
	}

	var leftP = 0
	var rightP = len(s) - 1
	var inputRune = []rune(s)
	for leftP < rightP {
		if !isVowel(inputRune[leftP]) {
			leftP++
		} else if !isVowel(inputRune[rightP]) {
			rightP--
		} else {
			inputRune[leftP], inputRune[rightP] = inputRune[rightP], inputRune[leftP]
			leftP++
			rightP--
		}
	}

	return string(inputRune)
}
