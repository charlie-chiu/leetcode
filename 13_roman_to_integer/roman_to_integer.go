package roman_to_integer

import (
	"strings"
)

var romanToInt = secondTry

// idea from leetcode discussion
// 4ms / 3.1MB
func secondTry(s string) int {
	romanValue := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var arabic int

	var inputLength = len(s)
	for i, char := range s {
		if i+1 < inputLength && romanValue[char] < romanValue[rune(s[i+1])] {
			arabic -= romanValue[char]
		} else {
			arabic += romanValue[char]
		}
	}

	return arabic
}

// 4ms / 3.1MB
func firstTry(s string) int {
	type Roman struct {
		Symbol string
		Value  int
	}
	var romans = []Roman{
		{"IV", 4},
		{"IX", 9},
		{"I", 1},
		{"V", 5},
		{"XL", 40},
		{"XC", 90},
		{"X", 10},
		{"L", 50},
		{"CD", 400},
		{"CM", 900},
		{"C", 100},
		{"D", 500},
		{"M", 1000},
	}

	var arabic int

	for len(s) > 0 {
		for _, roman := range romans {
			if strings.HasPrefix(s, roman.Symbol) {
				arabic += roman.Value
				s = strings.TrimPrefix(s, roman.Symbol)
			}
		}
	}

	return arabic
}
