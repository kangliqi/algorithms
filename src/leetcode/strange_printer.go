package leetcode

import "math"

// 664. 奇怪的打印机
// 有台奇怪的打印机有以下两个特殊要求：
// 	打印机每次只能打印由 同一个字符 组成的序列。
// 	每次可以在任意起始和结束位置打印新字符，并且会覆盖掉原来已有的字符。
// 给你一个字符串 s ，你的任务是计算这个打印机打印它需要的最少打印次数。
func strangePrinter(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
				continue
			}
			dp[i][j] = math.MaxInt64
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
			}
		}
	}
	return dp[0][n-1]
}
