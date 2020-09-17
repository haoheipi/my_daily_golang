package cn

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
// 示例 1:
//
// 输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//
//
// 示例 2:
//
// 输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//
//
// 示例 3:
//
// 输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
//
// Related Topics 哈希表 双指针 字符串 Sliding Window
// 👍 4322 👎 0

// Time: 2020-09-17 16:18:41

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	//m := make(map[byte]int)
	bitSet := [256]int{}

	maxLen, cur := 0, 1
	////pwwkew   s[p] = 0 s[w] = 1  2
	for i, l := 0, len(s); i < l; i++ {
		index := bitSet[s[i]]
		if index >= cur {
			maxLen = max(maxLen, i-cur+1)
			cur = index + 1
		}
		bitSet[s[i]] = i + 1
	}
	maxLen = max(maxLen, len(s)-cur+1)
	return maxLen
	//if len(s) == 0 {
	//	return 0
	//}
	//var bitSet [256]bool
	//result, left, right := 0, 0, 0
	//////pwwkew   s[p] = 0 s[w] = 1  2
	//for left < len(s) {
	//	// 右侧字符对应的 bitSet 被标记 true，说明此字符在 X 位置重复，需要左侧向前移动，直到将X标记为 false
	//	if bitSet[s[right]] {
	//		bitSet[s[left]] = false
	//		left++
	//	} else {
	//		bitSet[s[right]] = true
	//		right++
	//	}
	//	if result < right-left {
	//		result = right - left
	//	}
	//	if left+result >= len(s) || right >= len(s) {
	//		break
	//	}
	//}
	//return result
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//leetcode submit region end(Prohibit modification and deletion)
