// 125. Valid Palindrome
// A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.
// Given a string s, return true if it is a palindrome, or false otherwise.

// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.

// Example 2:
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.

// Example 3:
// Input: s = " "
// Output: true

// Explanation: s is an empty string "" after removing non-alphanumeric characters.
// Since an empty string reads the same forward and backward, it is a palindrome.
package main

import "unicode"

func validSymbol(r rune) bool {
	if '0' <= r && r <= '9' {
		return true
	}
	lower := unicode.ToLower(r)
	return 'a' <= lower && lower <= 'z'
}

func isPalindrome(s string) bool {
	ss := []rune(s)
	l := 0
	r := len(ss) - 1
	for l < r {
		if !validSymbol(ss[l]) {
			l++
			continue
		}
		if !validSymbol(ss[r]) {
			r--
			continue
		}
		if unicode.ToLower(ss[l]) != unicode.ToLower(ss[r]) {
			return false
		}
		r--
		l++
	}
	return true
}
