/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	count := make([]int, 2)
	for _, val := range bills {
		switch val {
		case 5:
			count[0]++
		case 10:
			if count[0] == 0 {
				return false
			}
			count[1]++
			count[0]--
		case 20:
			if count[1] >= 1 && count[0] >= 1 {
				count[0]--
				count[1]--
			} else if count[0] >= 3 {
				count[0] -= 3
			} else {
				return false
			}
		}
	}
	return true
}

// @lc code=end

