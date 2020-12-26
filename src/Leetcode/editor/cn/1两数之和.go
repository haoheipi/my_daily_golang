package cn

//给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//
// 你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
//
//
//
// 示例:
//
// 给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]
//
// Related Topics 数组 哈希表
// 👍 9122 👎 0

// Time: 2020-09-15 10:18:48

//Leetcode submit region begin(Prohibit modification and deletion)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		left := target - num
		if index, exit := m[left]; exit {
			return []int{i, index}
		} else {
			m[num] = i
		}
	}
	return nil
}

//Leetcode submit region end(Prohibit modification and deletion)
