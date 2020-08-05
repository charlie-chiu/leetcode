package reverse_string

var solution = recursive

func try(s []byte) {
	length := len(s)
	start, end := 0, length-1
	for start < length/2 {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// bad at performance?
func recursive(s []byte) {
	l := len(s)

	if l <= 1 {
		return
	}

	s[0], s[l-1] = s[l-1], s[0]
	recursive(s[1 : len(s)-1])
}
