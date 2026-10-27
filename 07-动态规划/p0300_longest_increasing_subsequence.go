//go:build p0041

// LeetCode 第300题 - 最长递增子序列 (Longest Increasing Subsequence)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、二分查找、动态规划
//
// 【题目描述】
// 给你一个整数数组 nums，找到其中最长严格递增子序列的长度。
// 子序列 是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。
// 例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的一个子序列。
//
// 【示例 1】
// 输入：nums = [10,9,2,5,3,7,101,18]
// 输出：4
// 解释：最长递增子序列是 [2,3,7,101]，因此长度为 4。
//
// 【示例 2】
// 输入：nums = [0,1,0,3,2,3]
// 输出：4
//
// 【示例 3】
// 输入：nums = [7,7,7,7,7,7,7]
// 输出：1
//
// 【数据范围】
// 1 <= nums.length <= 2500
// -10^4 <= nums[i] <= 10^4
//
// 【思路提示】基础解法 O(n^2)：dp[i] = 以 i 结尾的 LIS 长度，遍历 j<i 且 nums[j]<nums[i]。
// 进阶 O(n log n)：维护 tails 数组，用二分找第一个 >= nums[i] 的位置做替换。
//
// 题目链接：https://leetcode.cn/problems/longest-increasing-subsequence/
package main

// lengthOfLIS 返回最长严格递增子序列长度。
func lengthOfLIS(nums []int) int {
	// TODO: 实现
	return 0
}

func main() {
	lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
}
