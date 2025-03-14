// 200. Number of Islands
//
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

// Example 1:
// Input: grid = [
//   ["1","1","1","1","0"],
//   ["1","1","0","1","0"],
//   ["1","1","0","0","0"],
//   ["0","0","0","0","0"]
// ]
// Output: 1

// Example 2:
// Input: grid = [
//   ["1","1","0","0","0"],
//   ["1","1","0","0","0"],
//   ["0","0","1","0","0"],
//   ["0","0","0","1","1"]
// ]
// Output: 3

// Constraints:
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.
package main

func numIslands(grid [][]byte) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	var islands int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if !visited[i][j] && grid[i][j] == '1' {
				islands++
				dfs(i, j, grid, visited)
			}
		}
	}
	return islands
}

func dfs(fromM, fromN int, grid [][]byte, visited [][]bool) {
	if fromM < 0 || fromM >= len(grid) || fromN < 0 || fromN >= len(grid[0]) || visited[fromM][fromN] || grid[fromM][fromN] == '0' {
		return
	}
	visited[fromM][fromN] = true
	dfs(fromM-1, fromN, grid, visited)
	dfs(fromM+1, fromN, grid, visited)
	dfs(fromM, fromN-1, grid, visited)
	dfs(fromM, fromN+1, grid, visited)
}
