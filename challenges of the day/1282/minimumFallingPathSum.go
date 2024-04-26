package main

import (
	"fmt"
	"math"
)

func minFallingPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// Initialize the first row
	for i := range grid[0] {
		dp[0][i] = grid[0][i]
	}

	// Iterate through the grid from the second row
	for y := 1; y < m; y++ {
		for x := 0; x < n; x++ {
			// Find the minimum sum from the previous row excluding the cell above
			prevMin := math.MaxInt32
			for i := 0; i < n; i++ {
				if i != x {
					prevMin = min(prevMin, dp[y-1][i])
				}
			}
			dp[y][x] = grid[y][x] + prevMin
		}
	}

	// Find the minimum element in the last row
	minSum := math.MaxInt32
	for _, val := range dp[m-1] {
		minSum = min(minSum, val)
	}
	return minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(minFallingPathSum(grid))
}
