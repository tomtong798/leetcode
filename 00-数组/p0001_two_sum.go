//go:build p0041

// LeetCode 第1题 - 两数之和 (Two Sum)
//
// 【难度】简单 | 【面试频率】★★★★★ | 【标签】数组、哈希表
//
// 【题目描述】
// 给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值
// target 的那两个整数，并返回它们的数组下标。
// 你可以假设每种输入只会对应一个答案，并且你不能使用两次同样的元素。
// 你可以按任意顺序返回答案。
//
// 【示例 1】
// 输入：nums = [2,7,11,15], target = 9
// 输出：[0,1]
// 解释：nums[0] + nums[1] == 9，返回 [0, 1]。
//
// 【示例 2】
// 输入：nums = [3,2,4], target = 6
// 输出：[1,2]
//
// 【示例 3】
// 输入：nums = [3,3], target = 6
// 输出：[0,1]
//
// 【数据范围】
// 2 <= nums.length <= 10^4
// -10^9 <= nums[i] <= 10^9
// -10^9 <= target <= 10^9
// 只会存在一个有效答案
//
// 【思路提示】用哈希表存"数值->下标"，一次遍历：每到一个数先查 target-当前值 是否已出现过。
//
// 题目链接：https://leetcode.cn/problems/two-sum/
package main

import "fmt"

// twoSum 返回和为 target 的两个下标。
func twoSum(nums []int, target int) []int {
	var numsMap = make(map[int]int)
	for i, num := range nums {
		numsMap[num] = i
	}

	for i, num := range nums {
		v, ok := numsMap[target-num]
		if ok && v > i {
			return []int{i, v}
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}

// ──────────────────────────────────────────────────────────
// 【点评】
//
// ✅ 正确性：完全正确，`v > i` 巧妙避免了同一元素复用。
//
// ⚠️ 可改进：没必要分两遍遍历，一遍就够了。
//
//   - 当前做法：第一遍建完整 map，第二遍再查 complement。
//     即使答案就在前两个元素，也要先遍历完整个数组建 map。
//   - 更优做法：边遍历边查，遇到答案立即返回：
//
//     func twoSum(nums []int, target int) []int {
//         seen := make(map[int]int)
//         for i, num := range nums {
//             if j, ok := seen[target-num]; ok {
//                 return []int{j, i}
//             }
//             seen[num] = i
//         }
//         return nil
//     }
//
//   - 只需一次遍历，遇到答案立即返回
//   - 天然避免重复使用同一元素（查 complement 时当前元素还没入 map）
//   - 不需要 `v > i` 额外判断，代码更短、意图更清晰
//
// 📝 变量命名：`numsMap` 略冗余，`seen` 更能表达语义。
//
// 📊 总结：
//   正确性 ✅ | 时间 O(n) ✅ | 空间 O(n) ✅ | 简洁度可优化
// ──────────────────────────────────────────────────────────
