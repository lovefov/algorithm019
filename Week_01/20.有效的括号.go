/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
// 切题. 括号是成对出现的,如果把他存在容器中,有先进后出的特点
// 所以考虑用 filo 型的数据结构, 栈就很合适
// 整一个 hashmap 记录
func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 { //预处理,奇数长度一定不是,缩小问题空间
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := []byte{}
	for i := 0; i < n; i++ {
		if pairs[s[i]] > 0 { // 处理右括号,左括号不是 key 访问是返回 byte 空值 0
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}
			// 情况1 找到右括号 stack里面没有配对
			// 情况2 合法的序列不可能出现两个连续的右括号
			stack = stack[:len(stack)-1] //弹出匹配的左括号.
		} else { // 左括号入栈
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

// @lc code=end

