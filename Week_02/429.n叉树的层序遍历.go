/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */
//广度优先搜索,熟练默写即可.
var res [][]int

func levelOrder(root *Node) [][]int {
	res = [][]int{}
	if root == nil {
		return res
	}

	queue := []*Node{root}
	var level int
	for 0 < len(queue) {
		counter := len(queue)
		res = append(res, []int{})
		for i := 0; i < counter; i++ {
			if queue[i] != nil {
				res[level] = append(res[level], queue[i].Val)
				for _, n := range queue[i].Children {
					queue = append(queue, n)
				}
			}
		}
		queue = queue[counter:]
		level++
	}
	return res
}

// @lc code=end

