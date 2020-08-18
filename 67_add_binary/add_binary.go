package add_binary

import (
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	answerSlice := []int{}
	var getIntValue = func(str string, index int) int {
		if index < 0 {
			return 0
		}

		v, _ := strconv.Atoi(string(str[index]))
		return v
	}

	var indexA, indexB = len(a) - 1, len(b) - 1
	var valueA, valueB, carry int
	for {
		valueA = getIntValue(a, indexA)
		valueB = getIntValue(b, indexB)
		sum := (valueA + valueB + carry) % 2
		//log.Printf("previous carry:%d \n", carry)
		carry = (valueA + valueB + carry) / 2
		//log.Printf("a:%d b:%d sum:%d c:%d\n", valueA, valueB, sum, carry)
		answerSlice = append(answerSlice, sum)

		// end condition
		indexA--
		indexB--
		if indexA < 0 && indexB < 0 && carry == 0 {
			break
		}
	}

	//make answer string
	answerString := strings.Builder{}
	for i := len(answerSlice) - 1; i >= 0; i-- {
		answerString.WriteString(strconv.Itoa(answerSlice[i]))
	}

	return answerString.String()
}
