// 242. Valid Anagram
// Given two strings s and t, return true if t is an
// anagram
// of s, and false otherwise.

// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true

// Example 2:
// Input: s = "rat", t = "car"
// Output: false
package main

func isAnagram(s string, t string) bool {
	letters := [26]int{}
	for _, ch := range s {
		letters[ch-'a']++
	}
	for _, ch := range t {
		letters[ch-'a']--
	}
	for _, cnt := range letters {
		if cnt != 0 {
			return false
		}
	}
	return true
}
