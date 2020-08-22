package count_primes

import "log"

var solution = try3

//copy idea from leetcode
func try3(n int) int {
	var isComposite = make(map[int]bool, n)
	var count int

	for i := 2; i < n; i++ {
		if _, ok := isComposite[i]; !ok {
			count++

			for j := i * i; j < n; j += i {
				isComposite[j] = true
			}
		}
	}

	return count
}

// 40ms
func lc1(n int) int {
	if n < 3 {
		return 0
	}

	var ans int

	// Composite means it's made up of various elements, thus not prime
	// We assume all numbers from 1 -> n are not composite, or are all prime
	isComposite := make([]bool, n)
	for i := 2; i < n; i++ {
		if !isComposite[i] {
			ans++
		}
		for j := 2; j*i < n; j++ {
			// if a number can be made of i*j, then it is indeed a composite and it isn't prime
			isComposite[j*i] = true
		}
	}

	return ans
}

func try2(n int) int {
	var isComposite = make(map[int]bool, n)

	for i := 2; i*i < n; i++ {
		if _, ok := isComposite[i]; !ok {
			//isComposite[i] = true

			for j := i * i; j < n; j += i {
				isComposite[j] = true
			}
		}
	}

	//count not composite, start from 2
	var count int
	for i := 2; i < n; i++ {
		if _, ok := isComposite[i]; !ok {
			count++
		}
	}

	return count
}

// use sieve of eratosthenes to fine primes
func countPrimes(n int) int {
	var isPrime = make(map[int]bool, n)

	for i := 2; i*i < n; i++ {
		if _, ok := isPrime[i]; !ok {
			isPrime[i] = true

			for j := i * i; j < n; j += i {
				isPrime[j] = false
				log.Printf("i: %d, j: %d, isP: %v", i, j, isPrime)
			}
		}
	}

	//log.Println(isPrime)

	//count primes
	var count int
	for i := 0; i < n; i++ {
		if isPrime[i] {
			count++
		}
	}

	return count
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
