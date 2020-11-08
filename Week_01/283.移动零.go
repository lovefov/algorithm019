/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	lastfound := 0 //快速指针
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastfound] = nums[i]
			lastfound++
		}
	}
	for i := lastfound; i < len(nums); i++ {
		nums[i] = 0
	}
}

// @lc code=end

