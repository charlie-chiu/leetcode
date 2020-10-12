package fibonacci

// f0 = 0, f1 = 1
// return nth number in fibonacci sequence

//iterate
// benchmark with  20th number: 64.3 ns/op
// benchmark with 100th number:  249 ns/op
func fibonacci(n int) int {
	sequence := make([]int, n+1)

	for i := 0; i <= n; i++ {
		switch i {
		case 0:
			sequence[i] = 0
		case 1:
			sequence[i] = 1
		default:
			sequence[i] = sequence[i-1] + sequence[i-2]
		}
	}

	return sequence[n]
}

//recursive
// benchmark with 100th number: TLE
// benchmark with  20th number: 40942ns/op
//func fibonacci(n int) int {
//	if n == 0 {
//		return 0
//	}
//
//	if n == 1 {
//		return 1
//	}
//
//	return fibonacci(n-1) + fibonacci(n-2)
//}

// recursive with memorization
// benchmark with  20th number:  2135 ns/op
// benchmark with 100th number: 13971 ns/op
//func fibonacci(n int) int {
//	var fib = map[int]int{
//		0: 0,
//		1: 1,
//	}
//	return helper(n, fib)
//}
//func helper(n int, fib map[int]int) int {
//	if number, ok := fib[n]; ok {
//		return number
//	}
//
//	fib[n-1] = helper(n-1, fib)
//	fib[n-2] = helper(n-2, fib)
//
//	return fib[n-1] + fib[n-2]
//}
