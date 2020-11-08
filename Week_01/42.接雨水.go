/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	l, r := 0, len(height)-1
	lmax, rmax, ans := 0, 0, 0
	for l < r {
		if height[l] < height[r] {
			if height[l] >= lmax {
				lmax = height[l]
			} else {
				ans += lmax - height[l]
			}
			l++
		} else {
			if height[r] >= rmax {
				rmax = height[r]
			} else {
				ans += rmax - height[r]
			}
			r--
		}
	}
	return ans
}

// @lc code=end

