/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
// 通过从简单的情况 先穷举出 一些情况.
// 然后发现通项公式 状态转移方程
// 然后求解
// 最后一定是能拆解成类似的重复子问题.
func climbStairs(n int) int {
    p, q, r := 0, 0, 1
    for i := 1; i <= n; i++ {
        p = q
        q = r
        r = p + q
    }
    return r
}

// @lc code=end

