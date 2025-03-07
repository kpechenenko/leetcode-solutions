// 118. Pascal's Triangle

// Given an integer numRows, return the first numRows of Pascal's triangle.

// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:

// Example 1:
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

// Example 2:
// Input: numRows = 1
// Output: [[1]]

// Constraints:
// 1 <= numRows <= 30
package main

func generate(n int) [][]int {
	if n == 1 {
		return [][]int{{1}}
	}
	if n == 2 {
		return [][]int{{1}, {1, 1}}
	}
	res := make([][]int, n, n)
	res[0] = []int{1}
	res[1] = []int{1, 1}
	for i := 2; i < n; i++ {
		res[i] = make([]int, i+1, i+1)
		res[i][0] = 1
		res[i][len(res[i])-1] = 1
		for j := 1; j < len(res[i])-1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
