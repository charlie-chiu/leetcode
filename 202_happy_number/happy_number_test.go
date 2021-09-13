package happy_number

import (
	"log"
	"testing"
)

type TestCase struct {
	n      int
	answer bool
}

func TestIsHappy(t *testing.T) {
	var TestCases = []TestCase{
		//{1, true},
		//{2, false},
		{19, true},
	}

	for _, tc := range TestCases {
		t.Run("", func(t *testing.T) {
			got := isHappy(tc.n)
			if got != tc.answer {
				log.Println("wrong answer.")
				log.Printf("input: %d, want %v got %v", tc.n, tc.answer, got)
				t.Fail()
			}
		})
	}
}
