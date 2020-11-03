package group_anagrams

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input  []string
	answer [][]string
}

func TestGroupAnagrams(t *testing.T) {
	var TestCases = []TestCase{
		{[]string{""}, [][]string{
			{""},
		}},
		{[]string{"a"}, [][]string{
			{"a"},
		}},
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{
			{"eat", "tea", "ate"},
			{"tan", "nat"},
			{"bat"},
		}},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := groupAnagrams(tc.input)
			if !reflect.DeepEqual(got, tc.answer) {
				t.Logf("wrong answer")
				t.Logf("input: %v", tc.input)
				t.Logf(" want: %v", tc.answer)
				t.Logf("  got: %v", got)
				t.Fail()
			}
		})
	}
}
