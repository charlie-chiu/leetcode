package longest_common_prefix

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	answer := strings.Builder{}

OuterLoop:
	for i := 0; ; i++ {

		var currentChar byte
		for _, input := range strs {
			if i >= len(input) {
				break OuterLoop
			}

			if currentChar == 0 {
				currentChar = input[i]
				continue
			}

			if currentChar != input[i] {
				break OuterLoop
			}

		}

		answer.WriteByte(currentChar)
	}

	return answer.String()
}
