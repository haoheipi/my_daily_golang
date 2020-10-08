package cn

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
// 👍 1897 👎 0

// Time: 2020-10-08 15:27:49

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	// 空字符串直接返回 true
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, v := range s {
		if (v == '[') || (v == '(') || (v == '{') {
			stack = append(stack, v)
		} else if ((v == ']') && len(stack) > 0 && stack[len(stack)-1] == '[') ||
			((v == ')') && len(stack) > 0 && stack[len(stack)-1] == '(') ||
			((v == '}') && len(stack) > 0 && stack[len(stack)-1] == '{') {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
