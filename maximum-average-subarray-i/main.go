// 643. Maximum Average Subarray I

// Find a contiguous subarray whose length is equal to k that has the maximum average value and return this value. Any answer with a calculation error less than 10-5 will be accepted.

// Example 1:
// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000
// Explanation: Maximum average is (12 - 5 - 6 + 50) / 4 = 51 / 4 = 12.75

// Example 2:
// Input: nums = [5], k = 1
// Output: 5.00000

// Constraints:
// n == nums.length
// 1 <= k <= n <= 10^5
// -10^4 <= nums[i] <= 10^4
package main

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxAvg := float64(sum) / float64(k)
	for i := k; i < len(nums); i++ {
		sum -= nums[i-k]
		sum += nums[i]
		avg := float64(sum) / float64(k)
		if avg > maxAvg {
			maxAvg = avg
		}
	}
	return maxAvg
}
