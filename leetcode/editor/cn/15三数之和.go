package cn

import (
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复
//的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例：
//
// 给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为：
//[
//  [-1, 0, 1],
//  [-1, -1, 2]
//]
//
// Related Topics 数组 双指针
// 👍 2605 👎 0

// Time: 2020-09-23 14:50:04

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > 0 {
			break
		}
		third := n - 1
		second := i + 1
		target := -1 * nums[i]
		// 枚举 b -4 -1 -1 0 1 2
		// 需要保证 b 的指针在 c 的指针的左侧
		for second < third {
			// 需要和上一次枚举的数不相同
			if second > i+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			if nums[second]+nums[third] > target {
				third--
			} else if nums[second]+nums[third] < target {
				second++
			} else {
				res = append(res, []int{nums[i], nums[second], nums[third]})
				second++
			}
		}

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
