package implement_strstr

import (
	"strings"
)

var strStr = firstTry

func firstTry(haystack, needle string) int {
	if needle == "" {
		return 0
	}

	// won't into loop if needle has more char than haystack
	needleLength := len(needle)
	for i := needleLength; i <= len(haystack); i++ {
		if haystack[i-needleLength:i] == needle {
			return i - needleLength
		}
	}

	return -1
}

func goLibrary(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
