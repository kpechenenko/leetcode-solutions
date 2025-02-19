// 36. Valid Sudoku
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
// Note:

// A Sudoku board (partially filled) could be valid but is not necessarily solvable.
// Only the filled cells need to be validated according to the mentioned rules.

// Input: board =
// [["5","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: true
// Example 2:

// Input: board =
// [["8","3",".",".","7",".",".",".","."]
// ,["6",".",".","1","9","5",".",".","."]
// ,[".","9","8",".",".",".",".","6","."]
// ,["8",".",".",".","6",".",".",".","3"]
// ,["4",".",".","8",".","3",".",".","1"]
// ,["7",".",".",".","2",".",".",".","6"]
// ,[".","6",".",".",".",".","2","8","."]
// ,[".",".",".","4","1","9",".",".","5"]
// ,[".",".",".",".","8",".",".","7","9"]]
// Output: false
// Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

package main

import "strconv"

func isValidSudoku(board [][]byte) bool {
    rows := [10][10]bool{}
    cols := [10][10]bool{}
    squares := [3][3][10]bool{}
    const dot = byte('.')
    for i, row := range board { // O(9)
        for j, el := range row { // O(9)
            if dot == el {
                continue
            }
            digit, _ := strconv.Atoi(string(el))
            if rows[i][digit] {
                return false
            } else {
                rows[i][digit] = true
            }
            if cols[j][digit] {
                return false
            } else {
                cols[j][digit] = true
            }
            sqI := convertElementIdxToSquareIdx(i)
            sqJ := convertElementIdxToSquareIdx(j)
            if squares[sqI][sqJ][digit] {
                return false
            } else {
                squares[sqI][sqJ][digit] = true
            }

        }
    }   
    return true
}

func convertElementIdxToSquareIdx(i int) int {
    if 0 <= i && i < 3 {
        return 0
    }
    if 3 <= i && i < 6 {
        return 1
    }
    return 2
}