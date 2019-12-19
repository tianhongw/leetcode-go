package a

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1
	}

	return 1 + maxInt(maxDepth(root.left), maxDepth(root.right))
}

func maxInt(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
