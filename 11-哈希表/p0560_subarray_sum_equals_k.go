//go:build p0041

// LeetCode 第560题 - 和为 K 的子数组 (Subarray Sum Equals K)
//
// 【难度】中等 | 【面试频率】★★★★★ | 【标签】数组、哈希表、前缀和
//
// 【题目描述】
// 给你一个整数数组 nums 和一个整数 k，请你统计并返回该数组中和为 k 的连续子数组的个数。
// 子数组是数组中元素的连续非空序列。
//
// 【示例 1】
// 输入：nums = [1,1,1], k = 2
// 输出：2
//
// 【示例 2】
// 输入：nums = [1,2,3], k = 3
// 输出：2
//
// 【数据范围】
// 1 <= nums.length <= 2 * 10^4
// -1000 <= nums[i] <= 1000
// -10^7 <= k <= 10^7
//
// 【思路提示】前缀和 + 哈希表：维护当前前缀和 sum，
// 查 sum-k 在哈希表中出现次数并累加到答案，再把 sum 加入哈希表。
//
// 题目链接：https://leetcode.cn/problems/subarray-sum-equals-k/
package main

// subarraySum 返回和为 k 的连续子数组个数。
func subarraySum(nums []int, k int) int {
	// TODO: 实现
	return 0
}

func main() {
	subarraySum([]int{1, 1, 1}, 2)
	subarraySum([]int{1, 2, 3}, 3)
}
