package solve0104

/**
 * @index 104
 * @title 二叉树的最大深度
 * @difficulty 简单
 * @tags tree,depth-first-search,breadth-first-search,binary-tree
 * @draft false
 * @link https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
 * @frontendId 104
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

// maxDepth 递归版本
func maxDepth(root *TreeNode) int {
	treeMaxDepth := 0
	if root == nil {
		return treeMaxDepth
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	if leftDepth > rightDepth {
		treeMaxDepth = leftDepth + 1
	} else {
		treeMaxDepth = rightDepth + 1
	}
	return treeMaxDepth
}

// maxDepth_ 非递归版本
func maxDepth_(root *TreeNode) int {
	treeMaxDepth := 0
	if root == nil {
		return treeMaxDepth
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		levelLength := len(queue)
		for i := 0; i < levelLength; i++ {
			queueHead := queue[0]
			queue = queue[1:]
			if queueHead.Left != nil {
				queue = append(queue, queueHead.Left)
			}
			if queueHead.Right != nil {
				queue = append(queue, queueHead.Right)
			}
		}
		treeMaxDepth++
	}
	return treeMaxDepth
}
