// 283. Move Zeroes

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.

// Note that you must do this in-place without making a copy of the array.

// Example 1:

// Input: nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
// Example 2:

// Input: nums = [0]
// Output: [0]
package main

func moveZeroes(nums []int) {
	insertNonZeroIdx := 0
	for i, num := range nums { // O(N)
		if num != 0 {
			if insertNonZeroIdx != i {
				nums[insertNonZeroIdx] = nums[i]
				nums[i] = 0
			}
			insertNonZeroIdx++
		}
	}
}
