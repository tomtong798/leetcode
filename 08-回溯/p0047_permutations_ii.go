//go:build p0041

// LeetCode 第47题 - 全排列 II (Permutations II)
//
// 【难度】中等 | 【面试频率】★★★★ | 【标签】数组、回溯
//
// 【题目描述】
// 给定一个可包含重复数字的序列 nums，按任意顺序返回所有不重复的全排列。
//
// 【示例 1】
// 输入：nums = [1,1,2]
// 输出：[[1,1,2],[1,2,1],[2,1,1]]
//
// 【示例 2】
// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// 【数据范围】
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
// 【思路提示】排序 + used 数组，同层剪枝：i>0 && nums[i]==nums[i-1] && !used[i-1] 时跳过。
//
// 题目链接：https://leetcode.cn/problems/permutations-ii/
package main

// permuteUnique 返回所有不重复全排列。
func permuteUnique(nums []int) [][]int {
	// TODO: 实现
	return nil
}

func main() {
	permuteUnique([]int{1, 1, 2})
	permuteUnique([]int{1, 2, 3})
}
