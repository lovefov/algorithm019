/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
// 通过三次翻转完成目标
// 使用双指针翻转数组
func reverse(start, end int, arr []int) {
	h := start
	e := end
	for h < e {
		arr[h], arr[e] = arr[e], arr[h]
		h++
		e--
	}
}
func rotate(arr []int, s int) []int {
	n := len(arr)
	s %= n               //通过模运算处理掉过大的 s 值,就是 kn + s 这无意义的 kn 次.
	s = n - s            //不变的那一部分
	reverse(0, s-1, arr) // 反转0-2
	reverse(s, n-1, arr)
	reverse(0, n-1, arr)
	return arr
}

// @lc code=end

