/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	nodes := []*TreeNode{root}
	for len(nodes) > 0 {
		currentLevel := []int{}
		copy := nodes[:]
		nodes = []*TreeNode{}
		for _, node := range copy {
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		res = append(res, currentLevel)
	}
	return res
}

// @lc code=end

