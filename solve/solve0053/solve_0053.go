package solve0053

import "fmt"

/**
 * @index 53
 * @title 最大子数组和
 * @difficulty 简单
 * @tags array,divide-and-conquer,dynamic-programming
 * @draft false
 * @link https://leetcode-cn.com/problems/maximum-subarray/
 * @frontendId 53
 */

// nums = append(nums, -2, 1, -3, 4, -1, 2, 1, -5, 4)
func maxSubArray(nums []int) int {
	pre := 0
	maxSum := nums[0]
	for _, num := range nums {
		if pre > 0 {
			pre = pre + num
		} else {
			pre = num
		}
		if pre < maxSum {
			maxSum = pre
		}
		fmt.Println("pre", pre)
		fmt.Println("maxSum", maxSum)
	}
	return maxSum
}
