package count_segments

import "strings"

var countSegments = goLibrary

// 0ms / 1.9MB
func goLibrary(s string) int {
	return len(strings.Fields(s))
}

// 0ms / 2.1MB
func firstTry(s string) int {
	var segments = 0
	var inSegment = false
	for i, char := range s {
		if char != ' ' {
			inSegment = true
		}

		if inSegment && char == ' ' {
			segments++
			inSegment = false
		}

		if inSegment && i == len(s)-1 {
			segments++
		}
	}

	return segments
}
