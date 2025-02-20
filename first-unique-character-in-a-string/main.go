// 387. First Unique Character in a String
// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
// Example 1:
// Input: s = "leetcode"
// Output: 0

// Explanation:
// The character 'l' at index 0 is the first character that does not occur at any other index.

// Example 2:
// Input: s = "loveleetcode"
// Output: 2

// Example 3:
// Input: s = "aabb"
// Output: -1
package main

func firstUniqChar(s string) int {
	const numOfLettersInEnglishAlphabet = 26
	var freq [numOfLettersInEnglishAlphabet]int
	for _, ch := range s {
		freq[ch-'a']++
	}
	for i, ch := range s {
		if freq[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
