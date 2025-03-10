// 73. Set Matrix Zeroes

// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
// You must do it in place.

// Example 1:
// Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
// Output: [[1,0,1],[0,0,0],[1,0,1]]

// Example 2:
// Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
// Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]

// Constraints:
// m == matrix.length
// n == matrix[0].length
// 1 <= m, n <= 200
// -2^31 <= matrix[i][j] <= 2^31 - 1
package main

func setZeroes(matrix [][]int) {
	markedRows := make([]bool, len(matrix))
	markedCols := make([]bool, len(matrix[0]))
	for i, row := range matrix {
		for j, el := range row {
			if el == 0 {
				markedRows[i] = true
				markedCols[j] = true
			}
		}
	}
	for i, mark := range markedRows {
		if mark {
			for j := 0; j < len(matrix[i]); j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j, mark := range markedCols {
		if mark {
			for i := 0; i < len(matrix); i++ {
				matrix[i][j] = 0
			}
		}
	}
}
