/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

 func preorder(root *Node) []int {
    res := []int{}
    dfs(root, &res)
    return res
}

func dfs(root *Node, res *[]int) {
    if root != nil {
        *res = append(*res, root.Val)

        for _, node := range root.Children{
            dfs(node, res)
        }
    }
}
// @lc code=end

