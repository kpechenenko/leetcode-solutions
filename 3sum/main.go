// 15. 3Sum

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.

// Example 1:
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// Explanation:
// nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
// nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
// nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
// The distinct triplets are [-1,0,1] and [-1,-1,2].
// Notice that the order of the output and the order of the triplets does not matter.

// Example 2:
// Input: nums = [0,1,1]
// Output: []
// Explanation: The only possible triplet does not sum up to 0.

// Example 3:
// Input: nums = [0,0,0]
// Output: [[0,0,0]]
// Explanation: The only possible triplet sums up to 0.

// Constraints:
// 3 <= nums.length <= 3000
// -105 <= nums[i] <= 10^5
package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	fr := make(map[int]int)
	for _, num := range nums {
		fr[num]++
	}
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		fr[nums[i]]--
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			fr[nums[j]]--
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			target := 0 - nums[i] - nums[j]
			if cnt, ok := fr[target]; ok && cnt > 0 {
				res = append(res, []int{nums[i], nums[j], target})
			}
		}
		for j := i + 1; j < len(nums); j++ {
			fr[nums[j]]++
		}
	}
	return res
}
