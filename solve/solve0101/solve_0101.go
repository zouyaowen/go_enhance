package solve0101

/**
 * @index 101
 * @title 对称二叉树
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/symmetric-tree/
 * @frontendId 101
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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return checkSymmetric(root, root)
}

func checkSymmetric(Left *TreeNode, right *TreeNode) bool {
	if Left == nil && right == nil {
		return true
	}
	if Left == nil || right == nil {
		return false
	}
	return Left.Val == right.Val && checkSymmetric(Left.Left, right.Right) && checkSymmetric(Left.Right, right.Left)
}

func isSymmetric_(root *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, root, root)
	for len(queue) > 0 {
		left := queue[0]
		right := queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
