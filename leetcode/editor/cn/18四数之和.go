package cn

import "sort"

//给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c +
// d 的值与 target 相等？找出所有满足条件且不重复的四元组。
//
// 注意：
//
// 答案中不可以包含重复的四元组。
//
// 示例：
//
// 给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
//
//满足要求的四元组集合为：
//[
//  [-1,  0, 0, 1],
//  [-2, -1, 1, 2],
//  [-2,  0, 0, 2]
//]
//
// Related Topics 数组 哈希表 双指针
// 👍 645 👎 0

// Time: 2020-10-08 14:44:14

//leetcode submit region begin(Prohibit modification and deletion)
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := make([][]int, 0)
	for i := 0; i < n-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < n-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			a3 := j + 1
			a4 := n - 1
			// 枚举 b -4 -1 -1 0 1 2
			// 需要保证 b 的指针在 c 的指针的左侧
			for a3 < a4 {
				// 需要和上一次枚举的数不相同
				if a3 > j+1 && nums[a3] == nums[a3-1] {
					a3++
					continue
				}
				if nums[i]+nums[j]+nums[a3]+nums[a4] > target {
					a4--
				} else if nums[i]+nums[j]+nums[a3]+nums[a4] < target {
					a3++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[a3], nums[a4]})
					a3++
				}
			}

		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
