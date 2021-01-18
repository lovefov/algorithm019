/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start

func isValidSudoku(board [][]byte) bool {
	var hang [9]int       //行
	var lie [9]int        //列
	var hanglie [3][3]int //9个九宫格
	for index1, val1 := range board {
		for index2, _ := range val1 {
			if board[index1][index2] == '.' { //如果是.直接跳出
				continue
			}
			//用与运算判断行中是否出现某该元素，若出现直接返回false
			if hang[index1]&int(math.Pow(2, float64(board[index1][index2]-'0'))) == int(math.Pow(2, float64(board[index1][index2]-'0'))) {
				return false
			} else {
				//用或运算将该元素对应的位填充为1
				hang[index1] = hang[index1] | int(math.Pow(2, float64(board[index1][index2]-'0')))
			}
			//同上
			if lie[index2]&int(math.Pow(2, float64(board[index1][index2]-'0'))) == int(math.Pow(2, float64(board[index1][index2]-'0'))) {
				return false
			} else {
				lie[index2] = lie[index2] | int(math.Pow(2, float64(board[index1][index2]-'0')))
			}
			//同上
			if hanglie[index1/3][index2/3]&int(math.Pow(2, float64(board[index1][index2]-'0'))) == int(math.Pow(2, float64(board[index1][index2]-'0'))) {
				return false
			} else {
				hanglie[index1/3][index2/3] = hanglie[index1/3][index2/3] | int(math.Pow(2, float64(board[index1][index2]-'0')))
			}
		}
	}
	return true
}

// @lc code=end

