package solve0160

/**
 * @index 160
 * @title 相交链表
 * @difficulty 简单
 * @tags hash-table,linked-list,two-pointers
 * @draft false
 * @link https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
 * @frontendId 160
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

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	point1 := headA
	point2 := headB
	for point1 != nil && point2 != nil {
		if point1 == point2 {
			return point2
		}
		point1 = point1.Next
		point2 = point2.Next
		if point1 == nil && point2 != nil {
			point1 = headB
		}
		if point2 == nil && point1 != nil {
			point2 = headA
		}
	}
	return nil
}
