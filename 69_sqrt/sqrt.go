package sqrt

// ignore Newton's method ...tight now

func mySqrt(x int) int {
	var root = x / 2

	for root*root > x {
		root /= 2
	}

	for root*root <= x {
		root++
	}

	return root - 1
}

// built-in function got 0ms on leercode
//func mySqrt(x int) int {
//	return int(math.Sqrt(float64((x))))
//}

//first try
//func mySqrt(x int) int {
//	for root := 1 ;; root++ {
//		if root*root > x {
//			return root-1
//		}
//	}
//}
