package solve0001

/**
 * @index 1
 * @title 两数之和
 * @difficulty 简单
 * @tags array,hash-table
 * @draft false
 * @link https://leetcode-cn.com/problems/two-sum/
 * @frontendId 1
 */

func twoSum(nums []int, target int) []int {
	dataMap := make(map[int]int)
	for i, num := range nums {
		if v, ok := dataMap[target-num]; ok {
			return []int{v, i}
		}
		dataMap[num] = i
	}
	return []int{}
}
