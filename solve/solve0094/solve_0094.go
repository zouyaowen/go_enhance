package solve0094

/**
 * @index 94
 * @title 二叉树的中序遍历
 * @difficulty 简单
 * @tags stack,tree,depth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
 * @frontendId 94
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	if root.Left != nil {
		leftRes := inorderTraversal(root.Left)
		res = append(res, leftRes...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		rightRes := inorderTraversal(root.Right)
		res = append(res, rightRes...)
	}
	return res
}

// inorderTraversal_ 非递归版本
func inorderTraversal_(root *TreeNode) []int {
	res := make([]int, 0)
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return res
}
