package cn

import (
	"fmt"
	"math"
	"sort"
)

//给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和
//。假定每组输入只存在唯一答案。
//
//
//
// 示例：
//
// 输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//
//
//
//
// 提示：
//
//
// 3 <= nums.length <= 10^3
// -10^3 <= nums[i] <= 10^3
// -10^4 <= target <= 10^4
//
// Related Topics 数组 双指针
// 👍 579 👎 0

// Time: 2020-09-23 15:40:41

//leetcode submit region begin(Prohibit modification and deletion)
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	res := math.MaxInt64
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		third := n - 1
		second := i + 1
		// 枚举 b -4 -1 -1 0 1 2
		// 需要保证 b 的指针在 c 的指针的左侧
		for second < third {
			// 需要和上一次枚举的数不相同
			if second > i+1 && nums[second] == nums[second-1] {
				second++
				continue
			}
			tmp := nums[i] + nums[second] + nums[third]
			fmt.Println(tmp)
			if tmp > target {
				third--
			} else if tmp < target {
				second++
			} else {
				return target
			}
			if math.Abs(float64(target-tmp)) < math.Abs(float64(target-res)) {
				res = tmp
			}
		}

	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
