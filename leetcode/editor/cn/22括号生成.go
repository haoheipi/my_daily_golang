package cn

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例：
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
//
// Related Topics 字符串 回溯算法
// 👍 1355 👎 0

// Time: 2020-10-08 15:33:12

//leetcode submit region begin(Prohibit modification and deletion)
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	var res []string
	findGenerateParenthesis(n, n, "", &res)
	return res
}

func findGenerateParenthesis(lindex, rindex int, str string, res *[]string) {
	if lindex == 0 && rindex == 0 {
		*res = append(*res, str)
		return
	}
	if lindex > 0 {
		findGenerateParenthesis(lindex-1, rindex, str+"(", res)
	}
	if rindex > 0 && lindex < rindex {
		findGenerateParenthesis(lindex, rindex-1, str+")", res)
	}
}

//leetcode submit region end(Prohibit modification and deletion)
