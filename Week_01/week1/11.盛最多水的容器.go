/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	l, r := 0, len(height)-1
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	var maxValue int
	for l < r {
		area := min(height[l], height[r]) * (r - l)
		maxValue = max(maxValue, area)
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxValue
}

// @lc code=end

