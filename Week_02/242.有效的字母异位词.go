/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
// 明确需求, 是否包含大小写.是否有什么特殊字符,超过了 ASCII 范围
// 输入的长度
// 暴力法. 排序,比较两个 string ,不一样直接返 false
// 使用 hashmap ,能并行加速么? 实现一个线程安全的 hm
// 开两个 hm 比较是否一样. 开一个 hm 一个往里面加 一个减
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hm := make(map[rune]int, len(s))
	for _, v := range s {
		if _, ok := hm[v]; ok {
			hm[v]++
		} else {
			hm[v] = 1
		}
	}
	for _, val := range t {
		if _, ok := hm[val]; ok {
			hm[val]--
		} else {
			return false
		}
	}
	for _, val := range hm {
		if val != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

