/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	val := []int{}
	var preorder func(*TreeNode)
	preorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		val = append(val, n.Val)
		preorder(n.Left)
		preorder(n.Right)
	}
	preorder(root)
	return val
}

// @lc code=end

