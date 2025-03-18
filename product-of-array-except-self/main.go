// 238. Product of Array Except Self

// Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

// You must write an algorithm that runs in O(n) time and without using the division operation.

// Example 1:
// Input: nums = [1,2,3,4]
// Output: [24,12,8,6]

// Example 2:
// Input: nums = [-1,1,0,-3,3]
// Output: [0,0,9,0,0]

// Constraints:
// 2 <= nums.length <= 10^5
// -30 <= nums[i] <= 30
// The input is generated such that answer[i] is guaranteed to fit in a 32-bit integer.

// Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
package main

func productExceptSelf(nums []int) []int {
	prefix := make([]int, len(nums), len(nums))
	prefix[0] = 1
	for i := 1; i < len(prefix); i++ {
		prefix[i] = nums[i-1] * prefix[i-1]
	}
	suffix := make([]int, len(nums), len(nums))
	suffix[len(suffix)-1] = 1
	for i := len(suffix) - 2; i >= 0; i-- {
		suffix[i] = nums[i+1] * suffix[i+1]
	}
	product := make([]int, len(nums), len(nums))
	for i := 0; i < len(product); i++ {
		product[i] = prefix[i] * suffix[i]
	}
	return product
}
