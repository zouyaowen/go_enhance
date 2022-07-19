package solve0070

/**
 * @index 70
 * @title 爬楼梯
 * @difficulty 简单
 * @tags memoization,math,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/climbing-stairs/
 * @frontendId 70
 */

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	dp := make([]int, 100)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
