package solve0021

/**
 * @index 21
 * @title 合并两个有序链表
 * @difficulty 简单
 * @tags recursion,linked-list
 * @draft false
 * @link https://leetcode-cn.com/problems/merge-two-sorted-lists/
 * @frontendId 21
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

// mergeTwoLists 递归写法
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val <= list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

// mergeTwoLists_ 非递归写法
func mergeTwoLists_(list1 *ListNode, list2 *ListNode) *ListNode {
	preNode := &ListNode{}
	cur := preNode
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else {
		cur.Next = list2
	}
	return preNode.Next
}
