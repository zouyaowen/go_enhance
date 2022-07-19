package solve0136

/**
 * @index 136
 * @title 只出现一次的数字
 * @difficulty 简单
 * @tags bit-manipulation,array
 * @draft false
 * @link https://leetcode-cn.com/problems/single-number/
 * @frontendId 136
 */

func singleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res = res ^ num
	}
	return res
}
