// 7. Reverse Integer
//
// Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

// Example 1:
// Input: x = 123
// Output: 321

// Example 2:
// Input: x = -123
// Output: -321

// Example 3:
// Input: x = 120
// Output: 21

package main

import "math"

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}
	reversed := 0
	for x > 0 {
		lastDigit := x % 10
		reversed *= 10
		reversed += lastDigit
		x /= 10
	}
	reversed *= sign
	if reversed > math.MaxInt32 || reversed < math.MinInt32 {
		return 0
	}
	return reversed
}
