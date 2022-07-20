package solve0169

/**
 * @index 169
 * @title 多数元素
 * @difficulty 简单
 * @tags array,hash-table,divide-and-conquer,counting,sorting
 * @draft false
 * @link https://leetcode-cn.com/problems/majority-element/
 * @frontendId 169
 */

func majorityElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	countMap := make(map[int]int)
	for _, num := range nums {
		if v, ok := countMap[num]; ok {
			if v+1 > length/2 {
				return num
			}
			countMap[num] = v + 1
		} else {
			countMap[num] = 1
		}
	}
	panic("data err")
}
