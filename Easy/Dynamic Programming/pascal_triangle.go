package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		dp[i] = make([]int, i+1)

		dp[i][0] = 1
		dp[i][i] = 1

		for j := 1; j < i; j++ {
			dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
		}
	}

	return dp
}
