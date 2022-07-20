package solve0141

/**
 * @index 141
 * @title 环形链表
 * @difficulty 简单
 * @tags hash-table,linked-list,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/linked-list-cycle/
 * @frontendId 141
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 快慢指针方案
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	cur := head
	fastCur := head.Next
	for cur != nil && fastCur != nil {
		if cur == nil || fastCur == nil {
			return false
		}
		if cur == fastCur {
			return true
		}
		cur = cur.Next
		fastCur = fastCur.Next
		if fastCur == nil {
			return false
		}
		fastCur = fastCur.Next
	}
	return false
}
