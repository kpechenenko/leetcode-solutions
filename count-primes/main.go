// 204. Count Primes

// Given an integer n, return the number of prime numbers that are strictly less than n.

// Example 1:
// Input: n = 10
// Output: 4
// Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.

// Example 2:
// Input: n = 0
// Output: 0

// Example 3:
// Input: n = 1
// Output: 0

// Constraints:
// 0 <= n <= 5 * 10^6
package main

func countPrimes(n int) int {
	primes := make([]bool, n, n)
	for i := 2; i < len(primes); i++ {
		primes[i] = true
	}
	for divisor := 2; divisor < len(primes); divisor++ {
		if primes[divisor] {
			for i := divisor * 2; i < len(primes); i += divisor {
				primes[i] = false
			}
		}
	}
	cnt := 0
	for i := 0; i < len(primes); i++ {
		if primes[i] {
			cnt++
		}
	}
	return cnt
}
