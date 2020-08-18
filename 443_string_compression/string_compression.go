package string_compression

import (
	"strconv"
	"strings"
)

var compress = firstTry

/*
	Given an array of characters, compress it in-place.

	The length after compression must always be smaller than or equal to the original array.

	Every element of the array should be a character (not int) of length 1.

	After you are done modifying the input array in-place, return the new length of the array.


	Follow up:
	Could you solve it using only O(1) extra space?
*/

func firstTry(chars []byte) int {
	if len(chars) < 2 {
		return len(chars)
	}

	//two pointers from left, using additional space for compressed string
	var left, right, charCount int
	compressed := strings.Builder{}
	for right <= len(chars) {
		if right == len(chars) {
			charCount = right - left
			compressed.WriteByte(chars[left])
			if charCount > 1 {
				compressed.WriteString(strconv.Itoa(charCount))
			}
			break
		}

		if chars[left] != chars[right] {
			charCount = right - left
			compressed.WriteByte(chars[left])
			if charCount > 1 {
				compressed.WriteString(strconv.Itoa(charCount))
			}
			left = right
		}

		right++
	}

	copy(chars, compressed.String())

	return len(compressed.String())
}
