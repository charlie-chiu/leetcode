package palindrome

import (
	"regexp"
	"strings"

	"log"
)

var solution = firstTry

func firstTry(s string) bool {
	s = strings.ToLower(s)
	reg, err := regexp.Compile("[^a-zA-Z0-9]")
	if err != nil {
		log.Fatalf("regexp.compile error : %v", err)
	}
	s = reg.ReplaceAllString(s, "")

	split := strings.Split(s, "")
	left := 0
	right := len(s) - 1

	for left < right {
		if split[left] != split[right] {
			return false
		}
		left++
		right--
	}

	return true
}
