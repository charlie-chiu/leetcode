package integer_to_roman

import (
	"sort"
	"strings"
)

var intToRoman = secondTry

// 20ms / 3.4MB
func secondTry(num int) string {
	romans := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var answer = strings.Builder{}

	for num > 0 {
		//log.Println(num)
		for _, roman := range romans {
			if num >= roman.value {
				num -= roman.value
				answer.WriteString(roman.symbol)
				break
			}
		}
	}

	return answer.String()
}

//12ms / 5.8MB
func firstTry(num int) string {
	romanSymbols := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keys := make([]int, len(romanSymbols))
	for key := range romanSymbols {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	var answer = strings.Builder{}
	for num > 0 {
		for _, value := range keys {
			if num >= value {
				num -= value
				answer.WriteString(romanSymbols[value])
				break
			}
		}
	}

	return answer.String()
}
