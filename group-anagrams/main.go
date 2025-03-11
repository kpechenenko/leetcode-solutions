// 49. Group Anagrams
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.

// Example 1:
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

// Explanation:
// There is no string in strs that can be rearranged to form "bat".
// The strings "nat" and "tan" are anagrams as they can be rearranged to form each other.
// The strings "ate", "eat", and "tea" are anagrams as they can be rearranged to form each other.

// Example 2:
// Input: strs = [""]
// Output: [[""]]

// Example 3:
// Input: strs = ["a"]
// Output: [["a"]]

// Constraints:
// 1 <= strs.length <= 10^4
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.
package main

import "sort"

func generateKey(src string) string {
	s := []rune(src)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	return string(s)
}

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, s := range strs {
		k := generateKey(s)
		groups[k] = append(groups[k], s)
	}
	res := make([][]string, 0, len(groups))
	for k := range groups {
		res = append(res, groups[k])
	}
	return res
}
