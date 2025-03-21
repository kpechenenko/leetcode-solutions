// 628. Maximum Product of Three Numbers

// Given an integer array nums, find three numbers whose product is maximum and return the maximum product.

// Example 1:
// Input: nums = [1,2,3]
// Output: 6

// Example 2:
// Input: nums = [1,2,3,4]
// Output: 24

// Example 3:
// Input: nums = [-1,-2,-3]
// Output: -6
 
// Constraints:
// 3 <= nums.length <= 10^4
// -1000 <= nums[i] <= 1000
package main

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	p1 := nums[n-3] * nums[n-2] * nums[n-1]
	p2 := nums[0] * nums[1] * nums[n-1]
	if p1 > p2 {
		return p1
	}
	return p2
}
