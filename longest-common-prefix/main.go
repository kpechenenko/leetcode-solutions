// 14. Longest Common Prefix

// Write a function to find the longest common prefix string amongst an array of strings.

// If there is no common prefix, return an empty string "".

// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"

// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.

package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	prevPrefix := ""
	prefix := ""
	i := 0
	firstS := []rune(strs[0])
	for {
		if i >= len(firstS) {
			break
		}
		prevPrefix = prefix
		prefix += string(firstS[i])
		prefixValid := true
		for _, s := range strs {
			if !strings.HasPrefix(s, prefix) {
				prefixValid = false
				break
			}
		}
		if !prefixValid {
			prefix = prevPrefix
			break
		}
		i++
	}
	return prefix
}
