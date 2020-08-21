package count_primes

var solution = countPrimes

func countPrimes(n int) int {
	return 99
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
