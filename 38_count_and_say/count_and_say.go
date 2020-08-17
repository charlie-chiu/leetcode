package count_and_say

import (
	"strconv"
	"strings"
)

var countAndSay = recursive

//recursive
func recursive(n int) string {
	if n < 1 || n > 30 {
		return ""
	}

	if n == 1 {
		return "1"
	}

	str := recursive(n - 1)

	var charCount = 0
	var say = strings.Builder{}
	for i := 0; i < len(str); i++ {
		charCount++

		if i+1 >= len(str) || str[i+1] != str[i] {
			say.WriteString(strconv.Itoa(charCount))
			say.WriteString(string(str[i]))
			charCount = 0
		}
	}

	return say.String()
}

func firstTry(n int) string {
	if n < 1 || n > 30 {
		return ""
	}

	var say = func(str string) string {
		var charCount = 0
		var say = strings.Builder{}
		for i := 0; i < len(str); i++ {
			charCount++

			if i+1 >= len(str) || str[i+1] != str[i] {
				say.WriteString(strconv.Itoa(charCount))
				say.WriteString(string(str[i]))
				charCount = 0
			}
		}

		return say.String()
	}

	var answer = "1"
	for i := 1; i < n; i++ {
		answer = say(answer)
	}

	return answer
}
