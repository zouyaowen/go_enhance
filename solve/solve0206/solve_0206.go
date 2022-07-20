package solve0206

/**
 * @index 206
 * @title 反转链表
 * @difficulty 简单
 * @tags recursion,linked-list
 * @draft false
 * @link https://leetcode-cn.com/problems/reverse-linked-list/
 * @frontendId 206
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

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}
